// Dynamic base URL
const apiBase = window.location.origin;

// Helper for ws protocol
function wsProtocol() {
    return window.location.protocol === "https:" ? "wss:" : "ws:";
}
const wsBase = wsProtocol() + "//" + window.location.host;

// Fetch leaderboard
function fetchLeaderboard() {
    fetch(apiBase + "/leaderboard")
        .then(res => res.json())
        .then(data => {
            const list = document.getElementById("leaderboard");
            list.innerHTML = "";
            if (data.length === 0) {
                list.innerHTML = "<li>No data yet</li>";
            } else {
                data.forEach(item => {
                    const li = document.createElement("li");
                    li.className = "bar-container";
                    const label = document.createElement("span");
                    label.className = "bar-label";
                    label.innerText = `${item.rank}. ${item.username}`;
                    const scoreSpan = document.createElement("span");
                    scoreSpan.className = "score-value";
                    scoreSpan.innerText = item.score;
                    li.appendChild(label);
                    li.appendChild(scoreSpan);
                    list.appendChild(li);
                });
            }
        })
        .catch(err => {
            console.error("Fetch error:", err);
        });
}

// WebSocket
const ws = new WebSocket(wsBase + "/ws");

ws.onopen = () => {
    console.log("WebSocket connected");
};

ws.onmessage = (event) => {
    if (event.data === "update") {
        fetchLeaderboard();
    }
};

// Initial load
fetchLeaderboard();

// Submit form
const form = document.getElementById("scoreForm");
const usernameInput = document.getElementById("username");
const pressBtn = document.getElementById("pressBtn");

form.addEventListener("submit", function(e) {
    e.preventDefault();
    const username = usernameInput.value.trim();
    if (!username) return;
    pressBtn.disabled = true;

    const originalBtnContent = pressBtn.innerHTML;
    pressBtn.innerHTML = `<span class="loading-spinner" style="margin:0 auto;display:inline-block;vertical-align:middle;"></span>`;

    fetch(apiBase + "/submit", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, point: 1 })
    })
    .then(res => {
        if (!res.ok) return res.text().then(msg => { throw new Error(msg); });
        document.getElementById("formMessage").innerText = "";
        setTimeout(() => {
            pressBtn.disabled = false;
            pressBtn.innerHTML = originalBtnContent;
        }, 500);
    })
    .catch(err => {
        document.getElementById("formMessage").innerText = err.message || "Error submitting press";
        setTimeout(() => {
            pressBtn.disabled = false;
            pressBtn.innerHTML = originalBtnContent;
        }, 500);
    });
});