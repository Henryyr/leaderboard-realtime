function fetchLeaderboard() {
    fetch("http://localhost:3100/leaderboard")
        .then(res => res.json())
        .then(data => {
            const list = document.getElementById("leaderboard");
            list.innerHTML = "";
            if (data.length === 0) {
                list.innerHTML = "<li>No data yet</li>";
            } else {
                data.forEach(item => {
                    const li = document.createElement("li");
                    li.innerText = `${item.rank}. ${item.username} - ${item.score}`;
                    list.appendChild(li);
                });
            }
        })
        .catch(err => {
            console.error("Fetch error:", err);
        });
}

// WebSocket
const ws = new WebSocket("ws://localhost:3100/ws");

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
document.getElementById("scoreForm").addEventListener("submit", function(e) {
    e.preventDefault();
    const username = document.getElementById("username").value;
    const score = document.getElementById("score").value;
    fetch("http://localhost:3100/submit", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, score: Number(score) })
    })
    .then(res => res.text())
    .then(msg => {
        document.getElementById("formMessage").innerText = msg;
        document.getElementById("scoreForm").reset();
    })
    .catch(err => {
        document.getElementById("formMessage").innerText = "Error submitting score";
    });
});
