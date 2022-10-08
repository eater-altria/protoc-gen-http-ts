package util

import "testing"

func TestTransPascalToCamel(t *testing.T) {
	if ans := transPascalToCamel("HelloWorld"); ans != "helloWorld" {
		t.Errorf("HelloWorld transport to camelCase shold be helloWorld, but %s got", ans)
	}

	if ans := transPascalToCamel("Testing"); ans != "testing" {
		t.Errorf("Testing transport to camelCase shold be testing, but %s got", ans)
	}
}

func TestTransPascalToSnake(t *testing.T) {
	if ans := transPascalToSnake("HelloWorld"); ans != "hello_world" {
		t.Errorf("HelloWorld transport to snake_case shold be hello_world, but %s got", ans)
	}

	if ans := transPascalToSnake("Testing"); ans != "testing" {
		t.Errorf("Testing transport to snake_case shold be testing, but %s got", ans)
	}
}

func TestTransCamelToPascal(t *testing.T) {
	if ans := transCamelToPascal("helloWorld"); ans != "HelloWorld" {
		t.Errorf("helloWorld transport to PascalCase shold be HelloWorld, but %s got", ans)
	}

	if ans := transCamelToPascal("testing"); ans != "Testing" {
		t.Errorf("testing transport to PascalCase shold be Testing, but %s got", ans)
	}
}

func TestTransCamelToSnake(t *testing.T) {
	if ans := transCamelToSnake("helloWorld"); ans != "hello_world" {
		t.Errorf("HelloWorld transport to snake_case shold be hello_world, but %s got", ans)
	}

	if ans := transCamelToSnake("testing"); ans != "testing" {
		t.Errorf("testing transport to snake_case shold be testing, but %s got", ans)
	}
}

func TestTransSnakeToPascal(t *testing.T) {
	if ans := transSnakeToPascal("hello_world"); ans != "HelloWorld" {
		t.Errorf("hello_world transport to PascalCase shold be hello_world, but %s got", ans)
	}

	if ans := transSnakeToPascal("testing"); ans != "Testing" {
		t.Errorf("testing transport to PascalCase shold be Testing, but %s got", ans)
	}
}

func TestTransSnakeToCamel(t *testing.T) {
	if ans := transSnakeToCamel("hello_world"); ans != "helloWorld" {
		t.Errorf("hello_world transport to camelCase shold be helloWorld, but %s got", ans)
	}

	if ans := transSnakeToCamel("testing"); ans != "testing" {
		t.Errorf("testing transport to camelCase shold be testing, but %s got", ans)
	}
}

func TestTransformNameStyle(t *testing.T) {
	if ans, _ := TransformNameStyle("hello_world", PascalCase); ans != "HelloWorld" {
		t.Errorf("hello_world transport to PascalCase shold be HelloWorld, but %s got", ans)
	}
	if ans, _ := TransformNameStyle("hello_world", CamelCase); ans != "helloWorld" {
		t.Errorf("hello_world transport to camelCase shold be helloWorld, but %s got", ans)
	}

	if ans, _ := TransformNameStyle("hello_world", SnakeCase); ans != "hello_world" {
		t.Errorf("hello_world transport to snake_case shold be hello_world, but %s got", ans)
	}

	if ans, _ := TransformNameStyle("HelloWorld", PascalCase); ans != "HelloWorld" {
		t.Errorf("HelloWorld transport to PascalCase shold be HelloWorld, but %s got", ans)
	}
	if ans, _ := TransformNameStyle("HelloWorld", CamelCase); ans != "helloWorld" {
		t.Errorf("HelloWorld transport to camelCase shold be helloWorld, but %s got", ans)
	}

	if ans, _ := TransformNameStyle("HelloWorld", SnakeCase); ans != "hello_world" {
		t.Errorf("HelloWorld transport to snake_case shold be hello_world, but %s got", ans)
	}

	if ans, _ := TransformNameStyle("helloWorld", PascalCase); ans != "HelloWorld" {
		t.Errorf("helloWorld transport to PascalCase shold be HelloWorld, but %s got", ans)
	}
	if ans, _ := TransformNameStyle("helloWorld", CamelCase); ans != "helloWorld" {
		t.Errorf("helloWorld transport to camelCase shold be helloWorld, but %s got", ans)
	}

	if ans, _ := TransformNameStyle("helloWorld", SnakeCase); ans != "hello_world" {
		t.Errorf("helloWorld transport to snake_case shold be hello_world, but %s got", ans)
	}
}
