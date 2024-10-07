// Handle authentication state on page load
document.addEventListener('DOMContentLoaded', function() {
    const token = localStorage.getItem('token');
    const authButtons = document.getElementById('auth-buttons');
    const userInfo = document.getElementById('user-info');
    const usernameSpan = document.getElementById('username');
    const authSection = document.getElementById('auth-section');
    const analyzeSection = document.getElementById('analyze-section');

    if (token) {
        // Placeholder for username
        usernameSpan.textContent = 'User'; // Replace with actual username
        authButtons.style.display = 'none';
        userInfo.style.display = 'block';
        authSection.style.display = 'none';
        analyzeSection.style.display = 'block';
        loadResults();
        setInterval(loadResults, 5000); // Update results every 5 seconds
    } else {
        authButtons.style.display = 'block';
        userInfo.style.display = 'none';
        authSection.style.display = 'block';
        analyzeSection.style.display = 'none';
    }
});

// Logout function
function logout() {
    localStorage.removeItem('token');
    window.location.href = '/';
}

// Handle form submission for analyze bar
document.getElementById('analyze-form').addEventListener('submit', function(event) {
    event.preventDefault();
    const url = document.getElementById('website-url').value;
    const token = localStorage.getItem('token');

    fetch('/api/analyze', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify({ url: url })
    })
    .then(response => response.json())
    .then(data => {
        if (data.success) {
            loadResults();
        } else {
            alert('Failed to analyze website.');
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert('An error occurred while analyzing the website.');
    });
});

// Load analyzed results
function loadResults() {
    const token = localStorage.getItem('token');

    fetch('/api/results', {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${token}`
        }
    })
    .then(response => response.json())
    .then(data => {
        const resultsTable = document.getElementById('results-table');
        if (data.length === 0) {
            resultsTable.innerHTML = '<div class="notification is-warning">No results yet.</div>';
            return;
        }
        let html = `
            <table class="table is-fullwidth">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>URL</th>
                        <th>Status</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
        `;
        data.forEach(result => {
            html += `
                <tr>
                    <td>${result.id}</td>
                    <td>${result.url}</td>
                    <td>${result.status}</td>
                    <td>
                        <a class="button is-small is-info" href="/result_details.html?id=${result.id}" hx-get="/result_details.html?id=${result.id}" hx-target="#content" hx-swap="innerHTML">View Details</a>
                    </td>
                </tr>
            `;
        });
        html += `
                </tbody>
            </table>
        `;
        resultsTable.innerHTML = html;
    })
    .catch(error => {
        console.error('Error:', error);
        const resultsTable = document.getElementById('results-table');
        resultsTable.innerHTML = '<div class="notification is-danger">An error occurred while loading results.</div>';
    });
}