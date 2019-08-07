// Code generated by goa v3.0.3, DO NOT EDIT.
//
// books views
//
// Command:
// $ goa gen library/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// LibraryBooks is the viewed result type that is projected based on a view.
type LibraryBooks struct {
	// Type to project
	Projected *LibraryBooksView
	// View to render
	View string
}

// LibraryBooksView is a type that runs validations on a projected type.
type LibraryBooksView struct {
	Data []*BookView
}

// BookView is a type that runs validations on a projected type.
type BookView struct {
	ID         *int64
	Title      *string
	Annotation *string
	Author     *string
	// images are a list of book photos
	Images []string
	Status *string
}

var (
	// LibraryBooksMap is a map of attribute names in result type LibraryBooks
	// indexed by view name.
	LibraryBooksMap = map[string][]string{
		"default": []string{
			"data",
		},
	}
)

// ValidateLibraryBooks runs the validations defined on the viewed result type
// LibraryBooks.
func ValidateLibraryBooks(result *LibraryBooks) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateLibraryBooksView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateLibraryBooksView runs the validations defined on LibraryBooksView
// using the "default" view.
func ValidateLibraryBooksView(result *LibraryBooksView) (err error) {
	if result.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "result"))
	}
	return
}

// ValidateBookView runs the validations defined on BookView.
func ValidateBookView(result *BookView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "result"))
	}
	if result.Annotation == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("annotation", "result"))
	}
	if result.Author == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("author", "result"))
	}
	if result.Images == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("images", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	return
}
