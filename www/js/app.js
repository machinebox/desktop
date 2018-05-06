
document.getElementById('now').innerHTML = new Date()

setInterval(function(){
    window.external.invoke("something")
}, 1000)
