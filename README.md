#Go unit test

## Processus
1. Ecrire son fichier foo.go
2. Créer le fichier foo_test.go
3. Dans le fichier foo_test.go : le nom de la fonction permet de cibler une fonction dans foo.go
### Exemple
Dans foo.go :
```go
func myFunc() {
    //code
}
```
Dans foo_test.go :
```go
func testmyFunc() {
    //test code
}
```
Command pour lancer les tests dans foo_test.go
```sh
yayadmin:~ $ go test
```
