![github-banner](https://user-images.githubusercontent.com/96031819/232985479-1699bf81-7198-4172-b2a0-ecaba9cf0e64.png)
![donuts-are-good's followers](https://img.shields.io/github/followers/donuts-are-good?&color=555&style=for-the-badge&label=followers) ![donuts-are-good's stars](https://img.shields.io/github/stars/donuts-are-good?affiliations=OWNER%2CCOLLABORATOR&color=555&style=for-the-badge) ![donuts-are-good's visitors](https://komarev.com/ghpvc/?username=donuts-are-good&color=555555&style=for-the-badge&label=visitors)

# 🌈 Colors

A delightful Go package to add a splash of color to your console applications! With support for 24-bit hex colors, and text formatting options like bold, italic, and underline, you can easily brighten up your text output and give it more pizzazz!

## 📸 Screenshot 

![image](https://github.com/donuts-are-good/colors/assets/96031819/e483f1e2-d6f4-4123-b8d5-71148d611068)

## 🥳 Example

```go 
package main

import (
	"fmt"
	"github.com/donuts-are-good/colors"
)

func main() {
	fmt.Println(colors.Bold + colors.BrightGreen + "'Oh, I beg your pardon!'" + colors.NC)
	fmt.Println(colors.Italic + colors.BrightYellow + "cried Alice hastily, afraid that she had hurt the poor animal's feelings." + colors.NC)
	fmt.Println(colors.Underline + colors.BrightCyan + "She found herself in a long, low hall, which was lit up by a row of lamps hanging from the roof." + colors.NC)
	fmt.Println(colors.Magenta + "There were doors all round the hall," + colors.NC)
	fmt.Println(colors.Red + "but they were all locked; and when Alice had been all the way down one side and up the other," + colors.NC)
	fmt.Println(colors.BrightMagenta + "trying every door, she walked sadly down the middle, wondering how she was ever to get out again." + colors.NC)
	fmt.Println(colors.Yellow + "Suddenly she came upon a little three-legged table," + colors.NC)
	fmt.Println(colors.BrightBlue + "all made of solid glass;" + colors.NC)
	fmt.Println(colors.BrightRed + "there was nothing on it except a tiny golden key," + colors.NC)
	fmt.Println(colors.Green + "and Alice's first thought was that it might belong to one of the doors of the hall;" + colors.NC)
	fmt.Println(colors.BrightBlack + "but, alas! either the locks were too large, or the key was too small," + colors.NC)
	fmt.Println(colors.BrightPurple + "but at any rate it would not open any of them." + colors.NC)
	fmt.Println(colors.BrightWhite + "However, on the second time round, she came upon a low curtain" + colors.NC)
	fmt.Println(colors.Black + "she had not noticed before, and behind it was a little door about fifteen inches high:" + colors.NC)
	fmt.Println(colors.Cyan + "she tried the little golden key in the lock, and to her great delight it fitted!" + colors.NC)
	fmt.Println(colors.White + "Alice opened the door and found that it led into a small passage," + colors.NC)
	fmt.Println(colors.Purple + "not much larger than a rat-hole:" + colors.NC)
	fmt.Println(colors.Blue + "she knelt down and looked along the passage into the loveliest garden you ever saw." + colors.NC)
}

```
## 🍭 Available Colors and Text Formatting

**NOW!** Free eraser included! (`colors.NC`)

**`colors` supports standard ANSI as well as 24bit hex color!**

- Black  
- Red    
- Green  
- Yellow 
- Purple 
- Magenta
- Cyan   
- White  

**Text Formatting Options:**

- Bold
- Faint
- Italic
- Underline

**Background Colors Too!**

- BgBlack & BgBrightBlack
- BgRed & BgBrightRed
- BgGreen & BgBrightGreen
- BgYellow & BgBrightYellow
- BgBlue & BgBrightBlue
- BgMagenta & BgBrightMagenta
- BgCyan & BgBrightCyan
- BgWhite & BgBrightWhite

**24-bit Hex Color Support:**

Use the `colors.Hex()` function to convert a hex color code to a 24-bit ANSI escape code. For example:

```go
colorMe := "Hello my shoes are a beautiful " + colors.Hex("ffeebb") + "off-white" + colors.NC + " color and the laces are " + colors.Hex("ab2f93") + "darker."

fmt.Println(colorMe)
```

## 🎭 Operating System Support

No worries! `colors` works like a charm on all major operating systems (maybe more idk) Windows, macOS, and Linux.

## 💖 Contributing
If you find a bug or have a feature request, please open an issue on GitHub. Pull requests are also welcome.

## 💡 License
`colors` is licensed under the MIT license. Check out the LICENSE file for more info. If you're not into legal stuff, just enjoy the colors and have fun! 
