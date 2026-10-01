function get(key, fallback) {
    var v = window.hakari && window.hakari['game.' + key]
    return v !== undefined ? v : fallback
}

function render() {
    document.getElementById('name').textContent = get('playerName', '—')
    document.getElementById('hp').textContent = get('hp', '—')
    document.getElementById('level').textContent = get('level', '—')
}

window.hakariFileSelected = function (url, name) {
    if (!url) {
        console.error('hakari: file load failed')
        return
    }
    console.log('hakari: file selected:', name, url)
    var player = document.getElementById('player')
    if (player) {
        player.src = url
        player.play().catch(function (e) {
            console.warn('hakari: autoplay failed:', e)
        })
    }
}

render()
setInterval(render, 100)