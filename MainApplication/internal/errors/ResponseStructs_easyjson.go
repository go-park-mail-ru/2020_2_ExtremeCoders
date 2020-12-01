// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package errors

import (
	LetterModel "MainApplication/internal/Letter/LetterModel"
	UserModel "MainApplication/internal/User/UserModel"
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonE0ad3eebDecodeMainApplicationInternalErrors(in *jlexer.Lexer, out *LetterAns) {
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
		case "Code":
			out.Code = int(in.Int())
		case "Lid":
			out.Lid = uint64(in.Uint64())
		case "Letters":
			if in.IsNull() {
				in.Skip()
				out.Letters = nil
			} else {
				in.Delim('[')
				if out.Letters == nil {
					if !in.IsDelim(']') {
						out.Letters = make([]LetterModel.Letter, 0, 0)
					} else {
						out.Letters = []LetterModel.Letter{}
					}
				} else {
					out.Letters = (out.Letters)[:0]
				}
				for !in.IsDelim(']') {
					var v1 LetterModel.Letter
					easyjsonE0ad3eebDecodeMainApplicationInternalLetterLetterModel(in, &v1)
					out.Letters = append(out.Letters, v1)
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
func easyjsonE0ad3eebEncodeMainApplicationInternalErrors(out *jwriter.Writer, in LetterAns) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Code\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Code))
	}
	{
		const prefix string = ",\"Lid\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Lid))
	}
	{
		const prefix string = ",\"Letters\":"
		out.RawString(prefix)
		if in.Letters == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Letters {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjsonE0ad3eebEncodeMainApplicationInternalLetterLetterModel(out, v3)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LetterAns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE0ad3eebEncodeMainApplicationInternalErrors(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LetterAns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE0ad3eebEncodeMainApplicationInternalErrors(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LetterAns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE0ad3eebDecodeMainApplicationInternalErrors(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LetterAns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE0ad3eebDecodeMainApplicationInternalErrors(l, v)
}
func easyjsonE0ad3eebDecodeMainApplicationInternalLetterLetterModel(in *jlexer.Lexer, out *LetterModel.Letter) {
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
		case "Id":
			out.Id = uint64(in.Uint64())
		case "Sender":
			out.Sender = string(in.String())
		case "Receiver":
			out.Receiver = string(in.String())
		case "Theme":
			out.Theme = string(in.String())
		case "Text":
			out.Text = string(in.String())
		case "DateTime":
			out.DateTime = int64(in.Int64())
		case "IsWatched":
			out.IsWatched = bool(in.Bool())
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
func easyjsonE0ad3eebEncodeMainApplicationInternalLetterLetterModel(out *jwriter.Writer, in LetterModel.Letter) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"Sender\":"
		out.RawString(prefix)
		out.String(string(in.Sender))
	}
	{
		const prefix string = ",\"Receiver\":"
		out.RawString(prefix)
		out.String(string(in.Receiver))
	}
	{
		const prefix string = ",\"Theme\":"
		out.RawString(prefix)
		out.String(string(in.Theme))
	}
	{
		const prefix string = ",\"Text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"DateTime\":"
		out.RawString(prefix)
		out.Int64(int64(in.DateTime))
	}
	{
		const prefix string = ",\"IsWatched\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsWatched))
	}
	out.RawByte('}')
}
func easyjsonE0ad3eebDecodeMainApplicationInternalErrors1(in *jlexer.Lexer, out *AnswerGet) {
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
		case "Code":
			out.Code = uint16(in.Uint16())
		case "Description":
			out.Description = string(in.String())
		case "User":
			easyjsonE0ad3eebDecodeMainApplicationInternalUserUserModel(in, &out.User)
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
func easyjsonE0ad3eebEncodeMainApplicationInternalErrors1(out *jwriter.Writer, in AnswerGet) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Code\":"
		out.RawString(prefix[1:])
		out.Uint16(uint16(in.Code))
	}
	{
		const prefix string = ",\"Description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"User\":"
		out.RawString(prefix)
		easyjsonE0ad3eebEncodeMainApplicationInternalUserUserModel(out, in.User)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AnswerGet) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE0ad3eebEncodeMainApplicationInternalErrors1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AnswerGet) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE0ad3eebEncodeMainApplicationInternalErrors1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AnswerGet) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE0ad3eebDecodeMainApplicationInternalErrors1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AnswerGet) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE0ad3eebDecodeMainApplicationInternalErrors1(l, v)
}
func easyjsonE0ad3eebDecodeMainApplicationInternalUserUserModel(in *jlexer.Lexer, out *UserModel.User) {
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
		case "Id":
			out.Id = uint64(in.Uint64())
		case "Name":
			out.Name = string(in.String())
		case "Surname":
			out.Surname = string(in.String())
		case "Email":
			out.Email = string(in.String())
		case "Password":
			out.Password = string(in.String())
		case "Img":
			out.Img = string(in.String())
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
func easyjsonE0ad3eebEncodeMainApplicationInternalUserUserModel(out *jwriter.Writer, in UserModel.User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Surname\":"
		out.RawString(prefix)
		out.String(string(in.Surname))
	}
	{
		const prefix string = ",\"Email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"Password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	{
		const prefix string = ",\"Img\":"
		out.RawString(prefix)
		out.String(string(in.Img))
	}
	out.RawByte('}')
}