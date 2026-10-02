# Hakari

Легкий фреймворк для створення десктопних застосунків та ігор на Go з інтерфейсом на HTML, CSS та JavaScript.

Без CGo. Без Node.js. Один бінарник.

**З любов'ю з України** 🇺🇦

---

## Зміст

1. [Що таке Hakari](#розділ-1-що-таке-hakari)
2. [Встановлення](#розділ-2-встановлення)
3. [Екрани](#розділ-3-екрани)
4. [Прив'язка даних](#розділ-4-привʼязка-даних)
5. [Збереження / Завантаження](#розділ-5-збереження--завантаження)
6. [Файлові діалоги](#розділ-6-файлові-діалоги)
7. [Збірка ресурсів](#розділ-7-збірка-ресурсів)
8. [Створення релізу](#розділ-8-створення-релізу)
9. [Приклади](#розділ-9-приклади)

---

## Розділ 1. Що таке Hakari

**Hakari** — це легкий фреймворк для створення десктопних застосунків та ігор на **Go**, з інтерфейсом на **HTML, CSS та JavaScript**. Без CGo. Без Node.js. Один бінарник.

### Проблема

Go — чудова мова для бекенду, утиліт та CLI. Але коли справа доходить до **десктопних застосунків з нормальним інтерфейсом** — усе розвалюється.

Існуючі рішення:

- **Fyne** — чистий Go, але API незручний, потокобезпечність — ілюзія, а рендеринг далекий від нативного.
- **Wails** — гарний UI, але тягне за собою **Node.js** та **WebView2 SDK**, процес збірки складний, а бінарник роздутий.
- **Gio** — швидкий, але immediate-mode, крива навчання крута, і підходить не всім.
- **Electron** — взагалі не Go. 150+ МБ на пусте вікно.

Кожен з них вирішує **частину** проблеми, але створює **нову**. Хочеться: **завантажив → написав HTML → написав Go → запустив → працює**. Без залежностей, без збирачів, без танців з бубном.

**Цього не було. Тому я створив Hakari.**

### Ідея

Ідея прийшла **не одразу**. Я перепробував кілька GUI-фреймворків на Go. У кожного — свої проблеми: десь API незручний, десь продуктивність, десь залежності.

Одного разу я натрапив на **`webview`** — бібліотеку, яка надає **нативний WebView** без CGo, з prebuilt-бібліотеками. Вона працювала **з коробки**. І тоді я зрозумів: **Go може дати логіку, HTML може дати інтерфейс, а Hakari — зв'язати їх**.

Так з'явилася ідея: **фреймворк, який не воює з WebView, а обгортає його в людський API**.

Назва **Hakari** походить від персонажа з *Jujutsu Kaisen*, який жбурляє джекпотами. Це пасує ідеально: користувач шукає нормальний фреймворк — і тут йому **випадає джекпот**: **«ДЖЕКПОТ! Ось Hakari, мій фреймворк.»**

### Про мене

Мене звати **Михайло**. Я з **України** 🇺🇦, і мені **15 років**. Пишу на **C, C++, C#, Go, Python та Java** — не залипаю на одній мові; беру ту, що підходить під задачу.

Я — **любитель Java**. Це може звучати дивно для Go-проєкту — але Go привернув мене з тих самих причин, чому я люблю Java: **нативна компіляція під будь-яку платформу**, **рівномірний розподіл ресурсів** та **передбачувана продуктивність**. Go бере найкраще з філософії Java і йде далі.

До Hakari я створив кілька проєктів: **Wincess** — утиліта для Windows на C++ (обхід активаційних обмежень та персоналізація), **Owl-UI** — Python GUI фреймворк, **Live-Reload** — C++ DLL hot-swapper, та **багато іншого, що я вивчав і перевіряв на собі**.

Я не «професійний розробник». Я — **цікавий**. Я вчуся через створення. Я ламаю речі, щоб зрозуміти, як вони працюють. І я створюю інструменти, які хочу використовувати сам.

### Філософія

Hakari побудований на **п'яти принципах**:

1. **Простота.** Завантажив → запустив → працює. Без конфігів, без збирачів, без Node.js.
2. **Один бінарник.** Усі ассети вбудовані. Копіюєш `.exe` — усе працює.
3. **Прозорість.** DevTools як у браузері. Дебаж UI як веб-сторінку.
4. **Універсальність.** Не лише ігри. Будь-які UI-центричні застосунки.
5. **Швидкість.** Менше 100 МБ пам'яті. Швидка збірка. Жодних залежностей.

**Hakari — це не «ще один GUI-фреймворк».** Це **філософія**: Go для логіки, HTML для інтерфейсу, Hakari для зв'язку. Усе інше — твоя справа.

---

## Розділ 2. Встановлення

### Вимоги

- **Go 1.21+** — [завантажити тут](https://go.dev/dl/)
- Редактор коду — VS Code, GoLand або будь-який інший
- **Windows**: WebView2 Runtime (зазвичай уже встановлено на Windows 10/11)

### Встановлення Hakari

У папці проєкту:

```bash
go mod init myapp
go get github.com/Muchprow/hakari
```

### Перший проєкт

Створи `main.go`:

```go
package main

import (
	"embed"

	"github.com/Muchprow/hakari"
)

//go:embed assets
var assetsFS embed.FS

func main() {
	app := hakari.New("My App")

	app.Screen("main", func(s *hakari.Screen) {
		s.HTML("assets/screens/main/main.html")
	})

	if err := app.LoadHTMLFromFS(assetsFS, "assets/index.html"); err != nil {
		panic(err)
	}

	app.Start("main")
}
```

### Структура проєкту

```
myapp/
├── main.go
├── go.mod
├── go.sum
└── assets/
    ├── index.html
    ├── index.css
    ├── index.js
    └── screens/
        └── main/
            ├── main.html
            ├── main.css
            └── main.js
```

### `assets/index.html`

```html
<!DOCTYPE html>
<html lang="uk">
<head>
    <meta charset="UTF-8">
    <title>My App</title>
    <link rel="stylesheet" href="index.css">
</head>
<body>
    <div id="hakari-screen-container"></div>
    <script src="index.js"></script>
</body>
</html>
```

### `assets/screens/main/main.html`

```html
<h1>Привіт, Hakari!</h1>
```

### Запуск

```bash
go run main.go
```

Відкривається вікно з **«Привіт, Hakari!»** — ось і все. Без кроку збірки, без Node.js, без конфігів.

### Збірка бінарника

```bash
go build -ldflags="-H windowsgui" -o myapp.exe
```

`-H windowsgui` ховає консольне вікно на Windows. `.exe` самодостатній — копіюй куди завгодно, він працює.

---

## Розділ 3. Екрани

Кожен екран має власні **HTML, CSS та JS**. Екрани реєструються в Go і перемикаються через `GoTo`. Стан зберігається між переходами.

### Зареєструвати екран

```go
app.Screen("menu", func(s *hakari.Screen) {
	s.HTML("assets/screens/menu/menu.html")
})

app.Screen("game", func(s *hakari.Screen) {
	s.HTML("assets/screens/game/game.html")
})
```

### Запустити з екрану

```go
app.Start("menu")
```

### Перехід між екранами

```go
app.Screen("menu", func(s *hakari.Screen) {
	s.HTML("assets/screens/menu/menu.html")

	s.On("start", func() {
		s.GoTo("game")
	})
})
```

В HTML:

```html
<button onclick="window.menu_start()">Почати гру</button>
```

### Хуки життєвого циклу

```go
s.OnEnter(func() {
	// викликається щоразу, коли екран показується
})

s.OnLeave(func() {
	// викликається щоразу, коли екран ховається
})
```

### Чому стан зберігається

Коли ти переходиш на інший екран, DOM попереднього **ховається**, а не знищується. Таймери, JS-змінні та значення полів виживають. Коли ти повертаєшся — усе залишається як було.

---

## Розділ 4. Прив'язка даних

Go та JavaScript обмінюються даними через `window.hakari`. Використовуй `Set` для відправки, `Get` для читання, `On` для обробки подій.

### Відправити дані з Go

```go
s.Set("hp", 100)
s.Set("playerName", "Muchprow")
```

JavaScript читає:

```js
const hp = window.hakari['game.hp']
const name = window.hakari['game.playerName']
```

Примітка: дані мають простір імен за назвою екрану. `s.Set("hp", 100)` на екрані `game` стає `game.hp` в JS.

### Читати дані в Go

```go
hp := s.Get("hp").(int)
```

### Обробка подій з JS

```go
s.On("attack", func(dmg int) int {
	hp := s.Get("hp").(int) - dmg
	s.Set("hp", hp)
	return hp
})
```

В HTML:

```html
<button onclick="window.game_attack(10)">Атака</button>
```

В JS:

```js
await window.game_attack(10)
```

### Глобальні дані

Дані, встановлені на `app` (не `s`), доступні під власним ключем:

```go
app.Set("version", "1.0.0")
```

```js
window.hakari['version']
```

---

## Розділ 5. Збереження / Завантаження

Зберігай та завантажуй JSON-дані в стандартну папку конфігурації ОС.

### Куди пишуться файли

- **Windows:** `%APPDATA%\Hakari\<app-name>\`
- **Linux:** `~/.config/Hakari/<app-name>/`
- **macOS:** `~/Library/Application Support/Hakari/<app-name>/`

### Визначити структуру збереження

```go
type GameState struct {
	PlayerName string `json:"playerName"`
	HP         int    `json:"hp"`
	Level      int    `json:"level"`
}
```

### Зберегти

```go
state := GameState{
	PlayerName: "Muchprow",
	HP:         60,
	Level:      7,
}

if err := s.Save("autosave", state); err != nil {
	// обробити помилку
}
```

### Завантажити

```go
var state GameState
if err := s.Load("autosave", &state); err != nil {
	// обробити помилку
}
```

### Перевірити наявність збереження

```go
if s.SaveExists("autosave") {
	// завантажити
}
```

### Видалити

```go
s.Delete("autosave")
```

### Паттерн автосохранения

```go
s.OnEnter(func() {
	if s.Get("hp") == nil {
		var state GameState
		if s.SaveExists("autosave") {
			if err := s.Load("autosave", &state); err == nil {
				s.Set("playerName", state.PlayerName)
				s.Set("hp", state.HP)
				s.Set("level", state.Level)
				return
			}
		}
		s.Set("playerName", "Player")
		s.Set("hp", 100)
		s.Set("level", 1)
	}
})
```

---

## Розділ 6. Файлові діалоги

Відкривай будь-який файл з диска — аудіо, відео, зображення, документи — і використовуй його в UI.

### Відкрити файл

```go
s.On("pickFile", func() {
	s.OpenFile()
})
```

В HTML:

```html
<button onclick="window.game_pickFile()">Відкрити файл</button>
```

### Обробити вибраний файл

Користувач вибирає файл, Go зберігає його в кеш і повертає локальний URL. Підхоплюй через `window.hakariFileSelected`:

```js
window.hakariFileSelected = function (url, name) {
	if (!url) {
		console.error('hakari: file load failed')
		return
	}
	console.log('selected:', name, url)

	var player = document.getElementById('player')
	if (player) {
		player.src = url
	}
}
```

### Як це працює

1. Go відкриває прихований `<input type="file">`.
2. Користувач вибирає файл.
3. JS читає його як base64 і надсилає в Go.
4. Go зберігає його в `SaveDir/cache/` з ім'ям на основі SHA256.
5. Go повертає URL типу `http://127.0.0.1:54321/<hash>.mp3`.
6. JS використовує цей URL у `<audio>`, `<video>` або `<img>`.

### Чому локальний HTTP-сервер

Браузери блокують `file://` URL з міркувань безпеки. Локальний HTTP-сервер — стандартний обхідний шлях. Він слухає `127.0.0.1` на випадковому порту, віддає тільки файли з кеш-папки та захищає від path traversal.

---

## Розділ 7. Збірка ресурсів

Hakari вбудовує CSS та JS в один HTML-документ під час завантаження. Ти пишеш окремі файли; фреймворк їх об'єднує.

### Що ти пишеш

```
assets/screens/game/
├── game.html
├── game.css
└── game.js
```

### `game.html`

```html
<link rel="stylesheet" href="game.css">

<h1>Гра</h1>
<button onclick="window.game_attack()">Атака</button>

<script src="game.js"></script>
```

### Що отримує користувач

Один HTML-документ із вбудованими `<style>` та `<script>`. Жодних додаткових запитів, жодних `file://` шляхів, жодних CORS-проблем.

### Чому це важливо

- **Один бінарник.** Усі ассети вбудовані через `//go:embed`.
- **Без сервера.** Ніякого віддавання статики, ніякого розв'язування шляхів.
- **Звичний workflow.** Пиши HTML, CSS, JS як у вебі.

### Правила шляхів

Посилання на ассети всередині екрану розв'язуються відносно папки самого екрану. Тобто `game.css` у `game.html` означає `assets/screens/game/game.css`. Не потрібні довгі відносні шляхи.

---

## Розділ 8. Створення релізу

### Стандартна збірка

```bash
go build -o myapp.exe
```

Це створює робочий бінарник із прикріпленим консольним вікном.

### GUI-збірка (без консолі)

```bash
go build -ldflags="-H windowsgui" -o myapp.exe
```

`-H windowsgui` ховає консольне вікно на Windows. Рекомендовано для релізу.

### Кросплатформена збірка

Hakari **не використовує CGo**, тому кроскомпіляція працює з коробки:

```bash
GOOS=windows GOARCH=amd64 go build -o myapp.exe
GOOS=linux   GOARCH=amd64 go build -o myapp
GOOS=darwin  GOARCH=amd64 go build -o myapp
```

### Що всередині бінарника

Усе. `//go:embed` пакує всі ассети у виконуваний файл. Копіюй `.exe` на іншу машину — він працює без зовнішніх файлів.

### Розмір та пам'ять

- **Розмір бінарника:** зазвичай 5-15 МБ
- **Використання пам'яті:** менше 100 МБ для типових застосунків

---

## Розділ 9. Приклади

### Todo-список

Простий застосунок із додаванням, перемиканням та видаленням.

```go
type Todo struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

app.Screen("main", func(s *hakari.Screen) {
	s.HTML("assets/screens/main/main.html")

	s.On("add", func(text string) []Todo {
		todos := getTodos(s)
		todos = append(todos, Todo{Text: text})
		s.Set("todos", todos)
		return todos
	})

	s.On("toggle", func(index int) []Todo {
		todos := getTodos(s)
		todos[index].Done = !todos[index].Done
		s.Set("todos", todos)
		return todos
	})

	s.On("delete", func(index int) []Todo {
		todos := getTodos(s)
		todos = append(todos[:index], todos[index+1:]...)
		s.Set("todos", todos)
		return todos
	})
})

func getTodos(s *hakari.Screen) []Todo {
	raw := s.Get("todos")
	if raw == nil {
		return []Todo{}
	}
	return raw.([]Todo)
}
```

### Обгортка над сайтом

Відкрий зовнішній сайт у вікні Hakari.

```go
app.Screen("main", func(s *hakari.Screen) {
	s.HTML("assets/screens/main/main.html")
})

s.On("open", func() {
	s.GoTo("web")
})

app.Screen("web", func(s *hakari.Screen) {
	s.HTML("assets/screens/web/web.html")
})
```

У `web.html`:

```html
<iframe src="https://example.com" style="width:100%;height:100vh;border:none;"></iframe>
```

### Проста гра

Мінімальний клікер.

```go
app.Screen("game", func(s *hakari.Screen) {
	s.HTML("assets/screens/game/game.html")

	s.OnEnter(func() {
		if s.Get("score") == nil {
			s.Set("score", 0)
		}
	})

	s.On("click", func() int {
		score := s.Get("score").(int) + 1
		s.Set("score", score)
		return score
	})
})
```

У `game.html`:

```html
<h1>Рахунок: <span id="score">0</span></h1>
<button onclick="window.game_click()">Клік</button>

<script>
document.addEventListener('click', function () {
	var score = window.hakari['game.score']
	var el = document.getElementById('score')
	if (el && score !== undefined) el.textContent = score
})
</script>
```

### Більше прикладів

Повні робочі приклади живуть у папці [`examples/`](../examples) репозиторію.