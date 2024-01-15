var app = new Vue({
    el: '#app',
    data: {
        unordered: [],
        ordered: [],
        absent: [],
        uletter: '',
        oletter: '',
        oposition: 1,
        aletter: '',
        recommended: []
    },
    methods: {
        addUnordered() {
            const { uletter } = this
            if (uletter.length === 0) { return }
            this.unordered.push(uletter)
            this.uletter = ''
        },
        delUnordered() {
            this.unordered.pop()
        },
        addOrdered() {
            const { oletter, oposition } = this
            if (oletter.length === 0 || oposition < 1 || oposition > 5) { return }
            this.ordered.push({ position: oposition, letter: oletter })
            this.oletter = ''
            this.oposition = 1
        },
        delOrdered() {
            this.ordered.pop()
        },
        addAbsent() {
            const { aletter } = this
            if (aletter.length === 0) { return }
            this.absent.push(aletter)
            this.aletter = ''
        },
        delAbsent() {
            this.absent.pop()
        },
        apiRecommend() {
            let ord = {};
            this.ordered.forEach(el => {
                ord[el.position] = el.letter
            })
            this.recommended = ''
            const requestOptions = {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ unordered: this.unordered, ordered: ord, absent: this.absent })
            }
            fetch("api/v1/solver", requestOptions)
                .then(response => response.json())
                .then(data => (this.recommended = data))

        }
    }
})