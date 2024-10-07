// frontend/src/js/app.js

// Handle authentication state on page load
document.addEventListener('DOMContentLoaded', function() {
    const token = localStorage.getItem('token');
    const authButtons = document.getElementById('auth-buttons');
    const userInfo = document.getElementById('user-info');
    const usernameSpan = document.getElementById('username');

    if (token) {
        // Optionally decode the token to get the username
        // For simplicity, assuming the token contains the username in the payload
        // You can use a library like jwt-decode or handle it on the backend

        // Example using jwt-decode (if you include the library)
        /*
        const decoded = jwt_decode(token);
        usernameSpan.textContent = decoded.username;
        */

        // Placeholder for username
        usernameSpan.textContent = 'User'; // Replace with actual username
        authButtons.style.display = 'none';
        userInfo.style.display = 'block';
    } else {
        authButtons.style.display = 'block';
        userInfo.style.display = 'none';
    }
});

// Logout function
function logout() {
    localStorage.removeItem('token');
    window.location.href = '/';
}
