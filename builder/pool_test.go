package pool

import (
	"context"
	"testing"
	"time"
)

func Test_ItemFunctions(t *testing.T) {
	p := New(1)

	item, _ := p.Get(context.Background())

	symbol := "hello"

	_, _ = item.Write([]byte(symbol))

	if item.String() != symbol {
		t.Fatalf("Write() wanted %q, got %q", symbol, item.String())
	}

	item.Reset()

	_, _ = item.WriteString(symbol)

	if item.String() != symbol {
		t.Fatalf("Write() wanted %q, got %q", symbol, item.String())
	}

	item.Reset()

	runeValue := []rune(symbol)

	_, _ = item.WriteRune(runeValue[0])

	if item.String() != string(runeValue[0]) {
		t.Fatalf("Write() wanted %q, got %q", string(runeValue[0]), item.String())
	}

	item.Reset()

	_ = item.WriteByte(symbol[0])

	if item.String() != string(symbol[0]) {
		t.Fatalf("Write() wanted %q, got %q", string(symbol[0]), item.String())
	}
}

func Test_ItemNew(t *testing.T) {
	p := New(2)

	item, _ := p.Get(context.Background())

	if item.Index() != 0 {
		t.Fatalf("Index() wanted %d, got %d", 0, item.Index())
	}

	item.Close()

	item, _ = p.Get(context.Background())

	if item.Index() != 1 {
		t.Fatalf("Index() wanted %d, got %d", 1, item.Index())
	}
}

func Test_TimeoutCase(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond*100)
	defer cancel()

	p := New(1)

	_, err := p.Get(context.Background())
	if err != nil {
		t.Fatalf("Get(), wanted no error, got %v", err)
	}

	_, err = p.Get(ctx)
	if err == nil {
		t.Fatalf("Get() must be an ErrTimeoutDone")
	}
}

func Test_Validate(t *testing.T) {
	p := New(0)

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	_, err := p.Get(ctx)
	if err != nil {
		t.Fatal("Get() no error expected", err)
	}
}

func Test_Pooling(t *testing.T) {
	p := New(2)

	item, _ := p.Get(context.Background())

	if item.Index() != 0 {
		t.Fatalf("Index(), expected %d, got %d", 0, item.Index())
	}

	item.Close()

	item, _ = p.Get(context.Background())

	if item.Index() != 1 {
		t.Fatalf("Index(), expected %d, got %d", 1, item.Index())
	}

	item.Close()

	item, _ = p.Get(context.Background())

	if item.Index() != 0 {
		t.Fatalf("Index(), expected %d, got %d", 0, item.Index())
	}
}
