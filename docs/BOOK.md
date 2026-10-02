# Hakari

Lightweight Go framework for building UI-centric desktop apps and games with HTML, CSS and JavaScript.

No CGo. No Node.js. One binary.

**With love from Ukraine** 🇺🇦

---

## Table of Contents

1. [What is Hakari](#chapter-1-what-is-hakari)
2. [Installation](#chapter-2-installation)
3. [Screens](#chapter-3-screens)
4. [Data Binding](#chapter-4-data-binding)
5. [Save / Load](#chapter-5-save--load)
6. [File Dialogs](#chapter-6-file-dialogs)
7. [Asset Bundling](#chapter-7-asset-bundling)
8. [Building a Release](#chapter-8-building-a-release)
9. [Examples](#chapter-9-examples)
10. [About the Author](#chapter-10-about-the-author)

---

## Chapter 1. What is Hakari

**Hakari** is a lightweight framework for building desktop applications and games in **Go**, with a UI written in **HTML, CSS, and JavaScript**. No CGo. No Node.js. One binary.

### The Problem

Go is a great language for backends, utilities, and CLI tools. But when it comes to **desktop apps with a proper interface** — everything falls apart.

Existing solutions:

- **Fyne** — pure Go, but the API is awkward, thread safety is an illusion, and the rendering is far from native.
- **Wails** — beautiful UI, but pulls in **Node.js** and **WebView2 SDK**, build process is complicated, and the binary is bloated.
- **Gio** — fast, but immediate-mode, steep learning curve, and not for everyone.
- **Electron** — not even Go. 150+ MB for an empty window.

Each of them solves **part** of the problem but creates **a new one**. What you want is: **download → write HTML → write Go → run → it works**. No dependencies, no bundlers, no dancing with a tambourine.

**That didn't exist. So I built Hakari.**

### The Idea

The idea didn't come **all at once**. I tried several Go GUI frameworks. Each had its own problems — awkward APIs, poor performance, heavy dependencies.

Then I stumbled upon **`webview`** — a library that provides a **native WebView** without CGo, using prebuilt libraries. It **just worked**. And then I realized: **Go can handle the logic, HTML can handle the interface, and Hakari can connect them**.

That's how the idea was born: **a framework that doesn't fight WebView but wraps it in a human-friendly API**.

The name **Hakari** comes from a character in *Jujutsu Kaisen* who throws jackpots. It fits perfectly: a user is looking for a decent, **normal** framework — and then a jackpot hits them: **"JACKPOT! Here's Hakari, my framework."**

### About Me

My name is **Mykhailo** (Michael). I'm from **Ukraine** 🇺🇦, and I'm **15 years old**. I write in **C, C++, C#, Go, Java and Python** — I don't lock myself into one language; I pick whatever fits the task.

I'm a **Java enjoyer**. That might sound strange for a Go project — but Go caught my attention for the same reasons I love Java: **native compilation for any platform**, **even resource distribution**, and **predictable performance**. Go takes the best of Java's philosophy and pushes it further.

Before Hakari, I built several projects: **Wincess** — a Windows utility in C++ (bypassing activation watermarks and enabling personalization), **Owl-UI** — a Python GUI framework, **Live-Reload** — a C++ DLL hot-swapper, and **many other things I studied and tested on myself**.

I'm not a "professional developer." I'm a **curious one**. I learn by building. I break things to understand how they work. And I build tools I want to use myself.

### Philosophy

Hakari is built on **five principles**:

1. **Simplicity.** Download → run → it works. No configs, no bundlers, no Node.js.
2. **One binary.** All assets are embedded. Copy the `.exe` — everything works.
3. **Transparency.** DevTools like in a browser. Debug your UI as a web page.
4. **Versatility.** Not just games. Any UI-centric applications.
5. **Performance.** Under 100 MB of memory. Fast builds. No dependencies.

**Hakari is not "just another GUI framework."** It's a **philosophy**: Go for logic, HTML for the interface, Hakari for the connection. Everything else is up to you.

---

## Chapter 2. Installation

*Coming soon.*

---

## Chapter 3. Screens

*Coming soon.*

---

## Chapter 4. Data Binding

*Coming soon.*

---

## Chapter 5. Save / Load

*Coming soon.*

---

## Chapter 6. File Dialogs

*Coming soon.*

---

## Chapter 7. Asset Bundling

*Coming soon.*

---

## Chapter 8. Building a Release

*Coming soon.*

---

## Chapter 9. Examples

*Coming soon.*

---

## Chapter 10. About the Author

*Coming soon.*