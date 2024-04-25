package main

import (
	"encoding/binary"
	"io"
	"fmt"
	
)

type Payload interface{
	io.WriterTo
	io.ReaderFrom
	Bytes() []byte
}

type Binary []byte

func(m Binary) Bytes() []byte {
	return m
}

func (m Binary) WriteTo(w io.Writer) (int64,error){
	err := binary.Write(w, binary.BigEndian, uint32( len(m)))
	if err != nil{
		fmt.Println(err)
		return 0, err
	}
	n, err := w.Write(m)
	if err != nil{
		fmt.Println(err)
		return 0, err
	}
	return int64(n + 4), nil
}

func (m *Binary) ReadFrom(r io.Reader)(int64, error){
	
	var size uint32
	err := binary.Read(r, binary.BigEndian, &size)

	if err != nil{
		fmt.Println(err)
		return 0, err
	}

	*m = make([]byte, size)

	 n, err := r.Read(*m)
	 if err != nil{
		fmt.Println(err)
		return 0, err
	}
	return int64(n + 4), nil
}

func Decode(r io.Reader)(Payload, error){
	payload := new(Binary)
	_, err := payload.ReadFrom(r )
	if err != nil{
		fmt.Println(err)
		return nil, err
	}
	return payload, nil
}