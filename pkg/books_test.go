package booksapi

import (
	"context"
	books "library/gen/books"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

var testBooks = []*books.Book{
	&books.Book{ID: 1, Title: "Волшебник изумрудного города", Author: "А.С. Волков", Annotation: "Моя первая Книга", Images: []string{}},
	&books.Book{ID: 2, Title: "Хоббит. Туда и обратно", Author: "Толкиен", Annotation: "Так и не дочитал", Images: []string{}},
	&books.Book{ID: 3, Title: "Понедельник начинается в субботу!", Author: "Стругацкие", Annotation: "Привет, Стругацкие!", Images: []string{}},
}

type configurableBookRepo struct {
	getAllBooksFn      func() ([]*books.Book, error)
	getBookStatusFn    func(...int64) (map[int64]BookStatus, error)
	changeBookStatusFn func(bookID int64, subscriberID uuid.UUID, status string, createdBy uuid.UUID) error
}

func (r *configurableBookRepo) GetAllBooks() ([]*books.Book, error) {
	return r.getAllBooksFn()
}

func (r *configurableBookRepo) GetBookStatus(bookID ...int64) (map[int64]BookStatus, error) {
	return r.getBookStatusFn(bookID...)
}

func (r *configurableBookRepo) ChangeBookStatus(bookID int64, subscriberID uuid.UUID, status string, createdBy uuid.UUID) error {
	return r.changeBookStatusFn(bookID, subscriberID, status, createdBy)
}

func newConfgurableBooksRepo(
	GetBookStatusFn func(...int64) (map[int64]BookStatus, error),
	ChangeBookStatusFn func(bookID int64, subscriberID uuid.UUID, status string, createdBy uuid.UUID) error) *configurableBookRepo {
	r := configurableBookRepo{
		getAllBooksFn: func() ([]*books.Book, error) {
			return testBooks, nil
		},
		getBookStatusFn:    GetBookStatusFn,
		changeBookStatusFn: ChangeBookStatusFn,
	}

	return &r
}

func bookHasStatus(status string) func(...int64) (map[int64]BookStatus, error) {
	return func(ids ...int64) (map[int64]BookStatus, error) {
		res := make(map[int64]BookStatus)
		for _, v := range ids {
			res[v] = BookStatus{Status: status, Who: randomUUID()}
		}
		return res, nil
	}

}

func bookHasStatusAndSubscriberID(status string, subscriberID *uuid.UUID) func(...int64) (map[int64]BookStatus, error) {
	return func(ids ...int64) (map[int64]BookStatus, error) {
		res := make(map[int64]BookStatus)
		for _, v := range ids {
			res[v] = BookStatus{Status: status, Who: subscriberID, When: time.Now()}
		}
		return res, nil
	}
}

func alwaysReservedBook(_ int64) (string, *uuid.UUID, time.Time, error) {
	userID, _ := uuid.NewRandom()
	return statusFree, &userID, time.Now(), nil
}

func changeStatusOK(bookID int64, subscriberID uuid.UUID, status string, createdBy uuid.UUID) error {
	return nil
}

func randomUUID() *uuid.UUID {
	val, _ := uuid.NewRandom()
	return &val
}

func randomUUIDContext() context.Context {
	return uuidConext(*randomUUID())
}

func uuidConext(val uuid.UUID) context.Context {
	return context.WithValue(context.Background(), contextKeyUserID, val)
}

var logger = log.New(os.Stdout, "", 0)

func Test_bookssrvc_Reserve(t *testing.T) {

	type fields struct {
		logger *log.Logger
		repo   BooksRepo
	}
	type args struct {
		ctx context.Context
		p   *books.ReservePayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes interface{}
		wantErr bool
	}{
		{
			"can reserve if free",
			fields{logger, newConfgurableBooksRepo(bookHasStatus(statusFree), changeStatusOK)},
			args{
				ctx: randomUUIDContext(),
				p:   &books.ReservePayload{BookID: 1, SubscriberID: randomUUID().String(), Token: ""},
			},
			"",
			false,
		},
		{
			"can reserve if returned",
			fields{logger, newConfgurableBooksRepo(bookHasStatus(statusReturned), changeStatusOK)},
			args{
				ctx: randomUUIDContext(),
				p:   &books.ReservePayload{BookID: 1, SubscriberID: randomUUID().String(), Token: ""},
			},
			"",
			false,
		},
		{
			"can't reserve if reserved",
			fields{logger, newConfgurableBooksRepo(bookHasStatus(statusReserved), changeStatusOK)},
			args{
				ctx: randomUUIDContext(),
				p:   &books.ReservePayload{BookID: 1, SubscriberID: randomUUID().String(), Token: ""},
			},
			"",
			true,
		},
		{
			"can't reserve if pickedup",
			fields{logger, newConfgurableBooksRepo(bookHasStatus(statusPickedUp), changeStatusOK)},
			args{
				ctx: randomUUIDContext(),
				p:   &books.ReservePayload{BookID: 1, SubscriberID: randomUUID().String(), Token: ""},
			},
			"",
			true,
		},
		{
			"runtime error if unknown status",
			fields{logger, newConfgurableBooksRepo(bookHasStatus("newBookStatus"), changeStatusOK)},
			args{
				ctx: randomUUIDContext(),
				p:   &books.ReservePayload{BookID: 1, SubscriberID: randomUUID().String(), Token: ""},
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &bookssrvc{
				logger: tt.fields.logger,
				repo:   tt.fields.repo,
			}
			gotRes, err := s.Reserve(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("%v: error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("%v: res = %v, want %v", tt.name, gotRes, tt.wantRes)
			}
		})
	}
}

func Test_bookssrvc_Pickup(t *testing.T) {
	subscriberPickedUpID := randomUUID()

	type fields struct {
		logger *log.Logger
		repo   BooksRepo
	}
	type args struct {
		ctx context.Context
		p   *books.PickupPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes interface{}
		wantErr bool
	}{
		{
			"can pickup if reserved and the same user",
			fields{logger, newConfgurableBooksRepo(bookHasStatusAndSubscriberID("reserved", subscriberPickedUpID), changeStatusOK)},
			args{
				ctx: uuidConext(*subscriberPickedUpID),
				p:   &books.PickupPayload{BookID: 1, SubscriberID: subscriberPickedUpID.String()},
			},
			"",
			false,
		},
		{
			"can't pickup if reserved and not the same user",
			fields{logger, newConfgurableBooksRepo(bookHasStatus("reserved"), changeStatusOK)},
			args{
				ctx: randomUUIDContext(),
				p:   &books.PickupPayload{BookID: 1, SubscriberID: randomUUID().String()},
			},
			"",
			true,
		},
		{
			"can pickup if free",
			fields{logger, newConfgurableBooksRepo(bookHasStatus("free"), changeStatusOK)},
			args{
				ctx: uuidConext(*subscriberPickedUpID),
				p:   &books.PickupPayload{BookID: 1, SubscriberID: randomUUID().String()},
			},
			"",
			false,
		},
		{
			"can pickup if returned",
			fields{logger, newConfgurableBooksRepo(bookHasStatus("returned"), changeStatusOK)},
			args{
				ctx: uuidConext(*subscriberPickedUpID),
				p:   &books.PickupPayload{BookID: 1, SubscriberID: randomUUID().String()},
			},
			"",
			false,
		},
		{
			"runtime error if unknown status",
			fields{logger, newConfgurableBooksRepo(bookHasStatus("newBookStatus"), changeStatusOK)},
			args{
				ctx: randomUUIDContext(),
				p:   &books.PickupPayload{BookID: 1, SubscriberID: randomUUID().String()},
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &bookssrvc{
				logger: tt.fields.logger,
				repo:   tt.fields.repo,
			}
			gotRes, err := s.Pickup(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reserve(): error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Reserve(): res = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_bookssrvc_Return(t *testing.T) {

	subscriberPickedUpID := randomUUID()

	type fields struct {
		logger *log.Logger
		repo   BooksRepo
	}
	type args struct {
		ctx context.Context
		p   *books.ReturnPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes interface{}
		wantErr bool
	}{
		{
			"can return if reserved by the same user",
			fields{logger, newConfgurableBooksRepo(bookHasStatusAndSubscriberID(statusReserved, subscriberPickedUpID), changeStatusOK)},
			args{
				ctx: uuidConext(*subscriberPickedUpID),
				p:   &books.ReturnPayload{BookID: 1, SubscriberID: subscriberPickedUpID.String()},
			},
			"",
			false,
		},
		{
			"can return if pickedup by the same user",
			fields{logger, newConfgurableBooksRepo(bookHasStatusAndSubscriberID(statusPickedUp, subscriberPickedUpID), changeStatusOK)},
			args{
				ctx: uuidConext(*subscriberPickedUpID),
				p:   &books.ReturnPayload{BookID: 1, SubscriberID: subscriberPickedUpID.String()},
			},
			"",
			false,
		},
		{
			"can't return if reserved by another user",
			fields{logger, newConfgurableBooksRepo(bookHasStatusAndSubscriberID(statusReserved, subscriberPickedUpID), changeStatusOK)},
			args{
				ctx: uuidConext(*randomUUID()),
				p:   &books.ReturnPayload{BookID: 1, SubscriberID: randomUUID().String()},
			},
			"",
			true,
		},
		{
			"can't return if pickedup by another user",
			fields{logger, newConfgurableBooksRepo(bookHasStatusAndSubscriberID(statusPickedUp, subscriberPickedUpID), changeStatusOK)},
			args{
				ctx: uuidConext(*randomUUID()),
				p:   &books.ReturnPayload{BookID: 1, SubscriberID: randomUUID().String()},
			},
			"",
			true,
		},
		{
			"can't return if free",
			fields{logger, newConfgurableBooksRepo(bookHasStatus(statusFree), changeStatusOK)},
			args{
				ctx: uuidConext(*randomUUID()),
				p:   &books.ReturnPayload{BookID: 1, SubscriberID: randomUUID().String()},
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &bookssrvc{
				logger: tt.fields.logger,
				repo:   tt.fields.repo,
			}
			gotRes, err := s.Return(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("bookssrvc.Return() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("bookssrvc.Return() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
