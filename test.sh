echo "Ascii Art Justify Team"
go run . "Ascii Art Justify Team"
echo "Audit Questions"
echo "test 1:"
go run . --align=center "hello" standard  
echo "test 2:"
go run . --align=left "Hello There" standard 
echo "test 3:"
go run . --align=right "hello" shadow 
echo "test 4:"
go run . --align=justify "how are you" shadow 
echo "test 5:"
go run . --align=right left standard
echo "test 6:"
go run . --align=left right standard
echo "test 7:"
go run . --align=center hello shadow
echo "test 8:"
go run . --align=justify "1 Two 4" shadow
echo "test 9:"
go run . --align=right 23/32 standard
echo "test 10:"
go run . --align=right ABCabc123 thinkertoy
echo "test 11:"
go run . --align=center "#$%&\"" thinkertoy
echo "test 12:"
go run . --align=left '23Hello World!' standard
echo "test 13:"
go run . --align=justify 'HELLO there HOW are YOU?!' thinkertoy
echo "test 14:"
go run . --align=right "a -> A b -> B c -> C" shadow
echo "test 15:"
go run . --align=right abcd shadow
echo "test 16:"
go run . --align=center ola standard
echo "test 17:"
go run . --align=left "a -> A b -> B c -> C" shadow
echo "test 18:"
go run . --align=center "a -> A b -> B c -> C" thinkertoy
echo "test 19:"
go run . --align=justify "a -> A b -> B c -> C" standard
echo "test 20:"
go run . --align=justify "1 Two 4" standard
echo "test 21:"
go run . --align=hello "1 Two 4" standard