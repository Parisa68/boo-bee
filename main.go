package main

import (
	"fmt"
	"os/exec"
)

func main() {
	article := `Once, in a town surrounded by endless fields of sunflowers, there was a clockmaker named Felix. He was known not only for his skill but also for the magic he seemed to weave into every clock he made. His creations were always precise, but more than that, they had a curious tendency to affect time itself.

One day, a mysterious woman named Selene entered his shop, carrying a peculiar timepiece. It was an ancient pocket watch, its gold exterior dull and worn. "This watch," she said, her voice low and calm, "does not count time as it should. Can you fix it?"

Felix examined the watch, noticing something odd. Its hands moved backward slowly, as if trying to undo the past. He carefully opened the back, revealing intricate mechanisms that seemed almost alive, ticking and shifting on their own.

"I can try," Felix replied, intrigued. Over the next few days, he worked tirelessly on the watch, but every time he adjusted a gear or tightened a spring, the watch would reverse even faster. It seemed to resist any attempt to move forward.

Late one night, as Felix was on the verge of giving up, the watch suddenly stopped. For a moment, everything around him grew silent. The world outside his window froze—no wind, no rustling leaves, not even a sound from the busy street. Felix realized the watch wasn’t broken; it had been controlling time all along.

Selene returned the next morning. "It’s not meant to be fixed," Felix told her. "This watch... it controls time. I can feel it."

Selene smiled knowingly. "I didn’t ask you to fix it, Felix. I asked you to understand it."

In that moment, Felix felt the weight of the years lift off his shoulders. The clocks around him chimed all at once, marking a new beginning. Selene took the watch, leaving behind a single sunflower on his workbench, and with a nod, she disappeared as quietly as she had arrived.

From that day forward, Felix’s clocks never just told time—they gave people a chance to find moments they thought they had lost, offering glimpses into the past and a way to reshape their futures.`
	cmd := exec.Command("python3", "summarize.py", article)

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Summary:", string(out))
}
