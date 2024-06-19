### Inline-test

Репозиторий создан в целях проверки оптимизаций go компилятора по inline'у функций.

Были созданы модели (из http и слоя бизнес логики) и конверторы моделей между слоями.

Для чистоты эксперимента были созданы две версии конвертора, в одной из которых код намерянно усложнен путем добавления if statements.

Функции конвертаций размещены в отдельном пакете в отдельных файлах.

Некоторые конвертации принимают на вход модели по значению, некоторые по указателю.

Сбилдим на версии go 1.22.1:
```
➜  inline-test git:(main) ✗ go build -gcflags=-m main.go
# command-line-arguments
./main.go:58:31: inlining call to time.Time.UnixNano
./main.go:58:31: inlining call to time.(*Time).unixSec
./main.go:58:31: inlining call to time.(*Time).nsec
./main.go:58:31: inlining call to time.(*Time).sec

# нас интересует то что ниже, как видим упрощенная версия конверторов встроилась полностью
./main.go:33:34: inlining call to convert.ItemsToPBItems
./main.go:33:34: inlining call to convert.ItemToPBItem
./main.go:33:34: inlining call to convert.brandsToBrandsPB
./main.go:33:34: inlining call to convert.brandToBrandPB
```

Попробуем закомментировать if statements во второй версии конвертора `convert.ItemsToPBItems2` и сбилдим еще раз.
```go
		//if !item.IsActive || len(item.Title) < 100 {
		//	continue
		//}
```

```
➜  inline-test git:(main) ✗ go build -gcflags=-m main.go
# command-line-arguments
./main.go:58:31: inlining call to time.Time.UnixNano
./main.go:58:31: inlining call to time.(*Time).unixSec
./main.go:58:31: inlining call to time.(*Time).nsec
./main.go:58:31: inlining call to time.(*Time).sec
./main.go:33:34: inlining call to convert.ItemsToPBItems
./main.go:33:34: inlining call to convert.ItemToPBItem
./main.go:33:34: inlining call to convert.brandsToBrandsPB
./main.go:33:34: inlining call to convert.brandToBrandPB
./main.go:53:35: inlining call to convert.ItemsToPBItems2
./main.go:53:35: inlining call to convert.ItemToPBItem
./main.go:53:35: inlining call to convert.brandsToBrandsPB
./main.go:53:35: inlining call to convert.brandToBrandPB

```
Видно, что теперь заинлайнился и вызов функции `convert.ItemsToPBItems2` (и всех вложенных).

В версии 1.16 были добавлены некоторые улучшения по части inline'а функций.

> https://go.dev/doc/go1.16#compiler
> 
> The compiler can now inline functions with non-labeled for loops, method values, and type switches. The inliner can also detect more indirect calls where inlining is possible.

Однако здесь не указано про if statements.

В https://go.dev/wiki/CompilerOptimizations#function-inlining указано, что функции должны быть "simple enough" и не иметь сложных конструкций типа defer, select и т.д.
Возможно в следующих версиях инлайнер перестанет "спотыкаться" об if условия и сможет заинлайнить и с циклами и с условиями внутри циклов.