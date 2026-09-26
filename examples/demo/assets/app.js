function get(key, fallback) {
    return (window.hakari && window.hakari[key] !== undefined) ? window.hakari[key] : fallback
}

function render() {
    document.getElementById('name').textContent = get('playerName', '—')
    document.getElementById('hp').textContent = get('hp', '—')
    document.getElementById('level').textContent = get('level', '—')
}

async function attack() {
    await window.takeDamage(10)
    render()
}

async function up() {
    await window.levelUp()
    render()
}

async function init() {
    if (window.onReady) {
        await window.onReady()
    }
    render()
    setInterval(render, 100)
}

if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', init)
} else {
    init()
}