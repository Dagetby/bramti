package api

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/Dagetby/bramti/graph/model"
	"github.com/pkg/errors"
	"io"
	"strconv"
	"time"
)

type Twit struct {
	ID              int         `json:"id"`
	ContentText     string      `json:"contentText"`
	PublicationDate time.Time   `json:"publicationDate"`
	AuthorID        *model.User `json:"authorId"`
}

// Объявим базовый тип int для ID
func MarshalID(id int) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(fmt.Sprintf("%d", id)))
	})
}

// То же самое делается и при анмаршалинге
func UnmarshalID(v interface{}) (int, error) {
	id, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("ids must be strings")
	}
	i, e := strconv.Atoi(id)
	return int(i), e
}

func MarshalTimestamp(t time.Time) graphql.Marshaler {
	timestamp := t.Unix() * 1000

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(timestamp, 10))
	})
}

func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(int); ok {
		return time.Unix(int64(tmpStr), 0), nil
	}
	return time.Time{}, errors.New("Unmarshal Timestrap")
}
