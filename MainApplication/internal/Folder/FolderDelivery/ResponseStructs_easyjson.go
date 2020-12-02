// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package FolderDelivery

import (
	LetterModel "MainApplication/internal/Letter/LetterModel"
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

func easyjsonE0ad3eebDecodeMainApplicationInternalFolderFolderDelivery(in *jlexer.Lexer, out *SuccessAns) {
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
func easyjsonE0ad3eebEncodeMainApplicationInternalFolderFolderDelivery(out *jwriter.Writer, in SuccessAns) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Code\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Code))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SuccessAns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE0ad3eebEncodeMainApplicationInternalFolderFolderDelivery(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SuccessAns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE0ad3eebEncodeMainApplicationInternalFolderFolderDelivery(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SuccessAns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE0ad3eebDecodeMainApplicationInternalFolderFolderDelivery(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SuccessAns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE0ad3eebDecodeMainApplicationInternalFolderFolderDelivery(l, v)
}
func easyjsonE0ad3eebDecodeMainApplicationInternalFolderFolderDelivery1(in *jlexer.Lexer, out *LetterList) {
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
		case "Description":
			out.Description = string(in.String())
		case "Letter":
			if in.IsNull() {
				in.Skip()
				out.Letter = nil
			} else {
				in.Delim('[')
				if out.Letter == nil {
					if !in.IsDelim(']') {
						out.Letter = make([]LetterModel.Letter, 0, 0)
					} else {
						out.Letter = []LetterModel.Letter{}
					}
				} else {
					out.Letter = (out.Letter)[:0]
				}
				for !in.IsDelim(']') {
					var v1 LetterModel.Letter
					(v1).UnmarshalEasyJSON(in)
					out.Letter = append(out.Letter, v1)
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
func easyjsonE0ad3eebEncodeMainApplicationInternalFolderFolderDelivery1(out *jwriter.Writer, in LetterList) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Code\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Code))
	}
	{
		const prefix string = ",\"Description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"Letter\":"
		out.RawString(prefix)
		if in.Letter == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Letter {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LetterList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE0ad3eebEncodeMainApplicationInternalFolderFolderDelivery1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LetterList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE0ad3eebEncodeMainApplicationInternalFolderFolderDelivery1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LetterList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE0ad3eebDecodeMainApplicationInternalFolderFolderDelivery1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LetterList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE0ad3eebDecodeMainApplicationInternalFolderFolderDelivery1(l, v)
}
func easyjsonE0ad3eebDecodeMainApplicationInternalFolderFolderDelivery2(in *jlexer.Lexer, out *LetterErr) {
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
		case "Description":
			out.Description = string(in.String())
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
func easyjsonE0ad3eebEncodeMainApplicationInternalFolderFolderDelivery2(out *jwriter.Writer, in LetterErr) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Code\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Code))
	}
	{
		const prefix string = ",\"Description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LetterErr) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE0ad3eebEncodeMainApplicationInternalFolderFolderDelivery2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LetterErr) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE0ad3eebEncodeMainApplicationInternalFolderFolderDelivery2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LetterErr) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE0ad3eebDecodeMainApplicationInternalFolderFolderDelivery2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LetterErr) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE0ad3eebDecodeMainApplicationInternalFolderFolderDelivery2(l, v)
}
func easyjsonE0ad3eebDecodeMainApplicationInternalFolderFolderDelivery3(in *jlexer.Lexer, out *FolderList) {
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
		case "Folders":
			if in.IsNull() {
				in.Skip()
				out.Folders = nil
			} else {
				in.Delim('[')
				if out.Folders == nil {
					if !in.IsDelim(']') {
						out.Folders = make([]Folder, 0, 2)
					} else {
						out.Folders = []Folder{}
					}
				} else {
					out.Folders = (out.Folders)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Folder
					(v4).UnmarshalEasyJSON(in)
					out.Folders = append(out.Folders, v4)
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
func easyjsonE0ad3eebEncodeMainApplicationInternalFolderFolderDelivery3(out *jwriter.Writer, in FolderList) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Code\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Code))
	}
	{
		const prefix string = ",\"Folders\":"
		out.RawString(prefix)
		if in.Folders == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Folders {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FolderList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE0ad3eebEncodeMainApplicationInternalFolderFolderDelivery3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FolderList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE0ad3eebEncodeMainApplicationInternalFolderFolderDelivery3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FolderList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE0ad3eebDecodeMainApplicationInternalFolderFolderDelivery3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FolderList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE0ad3eebDecodeMainApplicationInternalFolderFolderDelivery3(l, v)
}
func easyjsonE0ad3eebDecodeMainApplicationInternalFolderFolderDelivery4(in *jlexer.Lexer, out *Folder) {
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
		case "Name":
			out.Name = string(in.String())
		case "Type":
			out.Type = string(in.String())
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
func easyjsonE0ad3eebEncodeMainApplicationInternalFolderFolderDelivery4(out *jwriter.Writer, in Folder) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Folder) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE0ad3eebEncodeMainApplicationInternalFolderFolderDelivery4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Folder) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE0ad3eebEncodeMainApplicationInternalFolderFolderDelivery4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Folder) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE0ad3eebDecodeMainApplicationInternalFolderFolderDelivery4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Folder) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE0ad3eebDecodeMainApplicationInternalFolderFolderDelivery4(l, v)
}
