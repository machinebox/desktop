
document.getElementById('now').innerHTML = new Date()

setInterval(function(){
    window.external.invoke("tick")
}, 1000)
