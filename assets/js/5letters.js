var app = new Vue({
    el: '#app',
    data: {
        words: [],
        recommended: [],
        wordInput: ''
    },
    methods: {
        addWord() {
            const { wordInput } = this
            let newWord = []
            if (wordInput.length !== 5) { return }
            for (const letter of wordInput.toLowerCase()) { newWord.push({ letter: letter, color: "grey" }) }
            this.words.push(newWord)
            this.wordInput = ''
        },
        delWord() {
            this.words.pop()
        },
        letterClicked(event) {
            coords = event.currentTarget.id.split("-")
            currentLetter = this.words[coords[0]][coords[1]].letter
            currentColor = this.words[coords[0]][coords[1]].color
            targetColor = ""
            if (currentColor == "grey") {
                targetColor = "white"
            }
            if (currentColor == "white") {
                targetColor = "yellow"
            }
            if (currentColor == "yellow") {
                targetColor = "grey"
            }
            this.words[coords[0]][coords[1]].color = targetColor
        },
        apiRecommend() {
            let unordered = {}
            let ordered = {}
            let absent = []
            let position = 0

            this.words.forEach(word => {
                position = 1
                word.forEach(letter => {
                    if (letter.color == "grey") {
                        absent.push(letter.letter)
                    }
                    if (letter.color == "white") {
                        unordered[position] = (letter.letter)
                    }
                    if (letter.color == "yellow") {
                        ordered[position] = letter.letter
                    }
                    position += 1
                })
            })
            this.recommended = ''
            const requestOptions = {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ unordered: unordered, ordered: ordered, absent: absent })
            }
            fetch("api/v1/solver", requestOptions)
                .then(response => response.json())
                .then(data => (this.recommended = data))

        }
    }
})