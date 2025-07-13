# 🧠 Zythex Language

Zythex — A custom futuristic programming language for logic-based programming and writing **Zythract smart contracts** on the Zytherion blockchain.

> 🔥 Made with ❤️ using Go + VS Code extension support.

---

## ✨ Key Features

- ✅ Backend Logic: variables (`zylet`), conditionals (`zyif`, `zelxe`), return (`zyturn`)
- ✅ Smart Contract DSL: use `zythract` and `function` to write contracts like in Solidity
- ✅ Print Statement: `zyth.print <...>` like `console.log`
- ✅ Custom Compiler: run `.zthx` files using the `zythx.exe` binary
- ✅ VS Code Extension: syntax highlighting, custom logo/icon, and runner
- ✅ Import Support: use library files with `import ./libs/...`

---

## 📦 Code Example

```zthx
import ./libs/zythvote.zthx

zythract SimpleStorage: 
zylet savedData = 821 

function set(value): 
savedData = value 
emit DataStored(value) 

function get(): 
zyth.print <storedData>
