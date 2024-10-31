
const {createApp, ref} = Vue;

const app = createApp({
    //template:/*html*/`
    //<h1>{{message}}</h1>
    //<p>{{author}}</p>
    //`,
    setup(){
        const message = ref("I'm message...")
        const author = ref('BBB')

        const changeText=()=>{

            message.value = 'setTimeout message'
            author.value = 'TTT'
        }

        return{
            message,
            author,
            changeText,
        }
    }
})

app.mount('#myApp')



