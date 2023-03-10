package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()
	ans, err := arith.Addition(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}
	require.Equal(t, int32(2), ans)
}

func TestSubtraction(t *testing.T) {
	arith := NewAdapter()
	ans, err := arith.Subtraction(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}
	require.Equal(t, int32(0), ans)
}

func TestMultiplication(t *testing.T) {
	arith := NewAdapter()
	ans, err := arith.Multiplication(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}
	require.Equal(t, int32(1), ans)
}

func TestDivision(t *testing.T) {
	arith := NewAdapter()
	ans, err := arith.Division(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}
	require.Equal(t, int32(1), ans)
}
