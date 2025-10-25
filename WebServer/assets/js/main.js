document.addEventListener('DOMContentLoaded', () => {
    const authContainer = document.getElementById('auth-container');
    const chirpContainer = document.getElementById('chirp-container');
    const feedContainer = document.getElementById('chirps-feed');

    const loginForm = document.getElementById('login-form');
    const registerForm = document.getElementById('register-form');
    const chirpForm = document.getElementById('chirp-form');

    // --- State Management ---
    let token = localStorage.getItem('authToken');

    function updateUIForAuthState() {
        if (token) {
            authContainer.classList.add('hidden');
            chirpContainer.classList.remove('hidden');
        } else {
            authContainer.classList.remove('hidden');
            chirpContainer.classList.add('hidden');
        }
    }

    // --- API Calls ---
    async function fetchChirps() {
        try {
            const response = await fetch('/api/chirps');
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const chirps = await response.json();
            renderChirps(chirps);
        } catch (error) {
            console.error("Could not fetch chirps:", error);
            feedContainer.innerHTML = '<p>Could not load chirps. Please try again later.</p>';
        }
    }

    async function handleLogin(e) {
        e.preventDefault();
        const email = document.getElementById('login-email').value;
        const password = document.getElementById('login-password').value;

        try {
            const response = await fetch('/api/login', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ email, password })
            });
            if (!response.ok) {
                throw new Error('Login failed');
            }
            const data = await response.json();
            token = data.token;
            localStorage.setItem('authToken', token);
            updateUIForAuthState();
            loginForm.reset();
        } catch (error) {
            console.error('Login error:', error);
            alert('Login failed. Please check your credentials.');
        }
    }

    async function handleRegister(e) {
        e.preventDefault();
        const email = document.getElementById('register-email').value;
        const password = document.getElementById('register-password').value;

        try {
            const response = await fetch('/api/users', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ email, password })
            });
            if (response.status !== 201) { // Assuming 201 Created for success
                throw new Error('Registration failed');
            }
            // Automatically log in the user after registration or prompt them to
            alert('Registration successful! Please log in.');
            registerForm.reset();
        } catch (error) {
            console.error('Registration error:', error);
            alert('Registration failed. Please try again.');
        }
    }

     async function handleChirpSubmit(e) {
        e.preventDefault();
        const body = document.getElementById('chirp-body').value;

        try {
            const response = await fetch('/api/chirps', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify({ body })
            });
            if (response.status !== 201) {
                throw new Error('Chirp submission failed');
            }
            chirpForm.reset();
            fetchChirps(); // Refresh the feed
        } catch (error) {
            console.error('Chirp submission error:', error);
            alert('Failed to post chirp. Please make sure you are logged in.');
        }
    }


    // --- Rendering ---
    function renderChirps(chirps) {
        feedContainer.innerHTML = ''; // Clear existing chirps
        // The API returns an object with numeric keys, so we convert it to an array
        const chirpsArray = Object.values(chirps);
        chirpsArray.sort((a, b) => a.id - b.id); // Sort by ID to maintain order
        
        for (const chirp of chirpsArray) {
            const chirpElement = document.createElement('div');
            chirpElement.className = 'chirp';
            
            const authorElement = document.createElement('p');
            authorElement.className = 'author';
            // We don't have author email in the chirp response, so this is a placeholder
            authorElement.textContent = `User ID: ${chirp.author_id}`; 
            
            const bodyElement = document.createElement('p');
            bodyElement.textContent = chirp.body;

            chirpElement.appendChild(authorElement);
            chirpElement.appendChild(bodyElement);
            feedContainer.appendChild(chirpElement);
        }
    }

    // --- Event Listeners ---
    loginForm.addEventListener('submit', handleLogin);
    registerForm.addEventListener('submit', handleRegister);
    chirpForm.addEventListener('submit', handleChirpSubmit);


    // --- Initial Load ---
    updateUIForAuthState();
    fetchChirps();
});
