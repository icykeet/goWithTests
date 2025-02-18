# H2 После установки бинарный файл pkgsite будет помещён в каталог $GOPATH/bin (или $HOME/go/bin, если GOPATH не настроен).
go install golang.org/x/pkgsite/cmd/pkgsite@latest

# H2 Добавление pkgsite в PATH: Убедитесь, что каталог, куда установился pkgsite, добавлен в вашу переменную окружения PATH. Например, если вы используете bash или zsh, добавьте следующую строку в ваш .bashrc, .zshrc или другой конфигурационный файл:
export PATH=$PATH:$(go env GOPATH)/bin

# H3 В зависимости от оболочки:
source ~/.bashrc  # или source ~/.zshrc

# H3 Чтобы открыть документацию для текущего каталога (например, Go-модуля), выполните:
pkgsite -open .

