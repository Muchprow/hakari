function get(key, fallback) {
    var v = window.hakari && window.hakari['game.' + key]
    return v !== undefined ? v : fallback
}

function render() {
    document.getElementById('name').textContent = get('playerName', '—')
    document.getElementById('hp').textContent = get('hp', '—')
    document.getElementById('level').textContent = get('level', '—')
}

render()
setInterval(render, 100)