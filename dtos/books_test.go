package dtos

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"goginproj/models"
)

func TestCreatedBookPagedResponse(t *testing.T) {
	books := []models.Book{
		{
			ID:     0,
			Title:  "abc",
			Author: "someone",
			Genre:  "text",
		},
	}
	type args struct {
		books    []models.Book
		page     int
		pageSize int
		count    int
	}
	tests := []struct {
		name string
		args args
		want gin.H
	}{
		{
			name: "test pagination return",
			args: args{

				books:    books,
				page:     1,
				pageSize: 1,
				count:    1,
			},
			want: gin.H{
				"books":        books,
				"current_page": 1,
				"page_size":    1,
				"total_books":  1,
				"total_page":   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreatedBookPagedResponse(tt.args.books, tt.args.page, tt.args.pageSize, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatedBookPagedResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
