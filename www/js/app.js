// app.js
// Replace this file with your own JavaScript front-end client.

// This stub allows the script to execute in a normal web browser too.
window.external.invoke = window.external.invoke || function(){
    console.info("window.external.invoke", arguments)
}

// every second, update the time and send a tick to the server.
setInterval(function(){
    window.external.invoke("tick")
    makeRequest('get', '/time', null, function(result, status, xhr) {
        document.getElementById('now').innerHTML = result
    })
}, 1000)

// makeRequest is a simple AJAX wrapper.
// It is not recommended for production.
function makeRequest(method, url, data, callback) {
    var xhr = new XMLHttpRequest()
    xhr.open(method, url, true)
    xhr.onreadystatechange = function(){
        if(xhr.readyState == XMLHttpRequest.DONE) {
            callback(xhr.response, xhr.status, xhr)
            return
        }
    }
    xhr.send(data)
}
