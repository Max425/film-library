// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package dto

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	"time"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson1a61c37dDecodeGithubComMax425FilmLibraryGitInternalHttpServerHandlerDto(in *jlexer.Lexer, out *Actor) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "gender":
			out.Gender = string(in.String())
		case "birth_date":
			if data := in.Raw(); in.Ok() {
				dt, err := time.Parse("2006-01-02", string(data[1:len(data)-1]))
				out.BirthDate = dt
				in.AddError(err)
			}
		case "films":
			if in.IsNull() {
				in.Skip()
				out.Films = nil
			} else {
				in.Delim('[')
				if out.Films == nil {
					if !in.IsDelim(']') {
						out.Films = make([]*Film, 0, 8)
					} else {
						out.Films = []*Film{}
					}
				} else {
					out.Films = (out.Films)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Film
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Film)
						}
						easyjson1a61c37dDecodeGithubComMax425FilmLibraryGitInternalHttpServerHandlerDto1(in, v1)
					}
					out.Films = append(out.Films, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1a61c37dEncodeGithubComMax425FilmLibraryGitInternalHttpServerHandlerDto(out *jwriter.Writer, in Actor) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"gender\":"
		out.RawString(prefix)
		out.String(string(in.Gender))
	}
	{
		const prefix string = ",\"birth_date\":"
		out.RawString(prefix)
		out.Raw((in.BirthDate).MarshalJSON())
	}
	{
		const prefix string = ",\"films\":"
		out.RawString(prefix)
		if in.Films == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Films {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjson1a61c37dEncodeGithubComMax425FilmLibraryGitInternalHttpServerHandlerDto1(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Actor) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1a61c37dEncodeGithubComMax425FilmLibraryGitInternalHttpServerHandlerDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Actor) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1a61c37dEncodeGithubComMax425FilmLibraryGitInternalHttpServerHandlerDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Actor) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1a61c37dDecodeGithubComMax425FilmLibraryGitInternalHttpServerHandlerDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Actor) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1a61c37dDecodeGithubComMax425FilmLibraryGitInternalHttpServerHandlerDto(l, v)
}
func easyjson1a61c37dDecodeGithubComMax425FilmLibraryGitInternalHttpServerHandlerDto1(in *jlexer.Lexer, out *Film) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "release_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ReleaseDate).UnmarshalJSON(data))
			}
		case "rating":
			out.Rating = float64(in.Float64())
		case "actors":
			if in.IsNull() {
				in.Skip()
				out.Actors = nil
			} else {
				in.Delim('[')
				if out.Actors == nil {
					if !in.IsDelim(']') {
						out.Actors = make([]*Actor, 0, 8)
					} else {
						out.Actors = []*Actor{}
					}
				} else {
					out.Actors = (out.Actors)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *Actor
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(Actor)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Actors = append(out.Actors, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1a61c37dEncodeGithubComMax425FilmLibraryGitInternalHttpServerHandlerDto1(out *jwriter.Writer, in Film) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"release_date\":"
		out.RawString(prefix)
		out.Raw((in.ReleaseDate).MarshalJSON())
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.Float64(float64(in.Rating))
	}
	{
		const prefix string = ",\"actors\":"
		out.RawString(prefix)
		if in.Actors == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Actors {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
