let token = localStorage.getItem('token');


document.getElementById('show-register').addEventListener('click', (e) => {
    e.preventDefault();
    document.getElementById('auth-form').style.display = 'none';
    document.getElementById('register-form').style.display = 'block';
});

document.getElementById('show-login').addEventListener('click', (e) => {
    e.preventDefault();
    document.getElementById('register-form').style.display = 'none';
    document.getElementById('auth-form').style.display = 'block';
});


document.getElementById('login-form').addEventListener('submit', async (e) => {
    e.preventDefault();
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    try {
        const response = await fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password })
        });

        if (!response.ok) throw new Error('Ошибка авторизации');

        const data = await response.json();
        token = data.token;
        localStorage.setItem('token', token);
        showTasksInterface();
        loadTasks();
    } catch (error) {
        alert(error.message);
    }
});


document.getElementById('register-form').addEventListener('submit', async (e) => {
    e.preventDefault();
    const email = document.getElementById('reg-email').value;
    const password = document.getElementById('reg-password').value;

    try {
        const response = await fetch('http://localhost:8080/register', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password })
        });

        if (!response.ok) throw new Error('Ошибка регистрации');

        alert('Регистрация успешна! Теперь войдите.');
        document.getElementById('register-form').style.display = 'none';
        document.getElementById('auth-form').style.display = 'block';
    } catch (error) {
        alert(error.message);
    }
});


document.getElementById('logout-btn').addEventListener('click', () => {
    localStorage.removeItem('token');
    token = null;
    document.getElementById('tasks-interface').style.display = 'none';
    document.getElementById('auth-form').style.display = 'block';
});


async function loadTasks() {
    if (!token) return;

    try {
        const response = await fetch('http://localhost:8080/tasks', {
            headers: { 'Authorization': token }
        });

        if (!response.ok) throw new Error('Ошибка загрузки задач');

        const tasks = await response.json();
        renderTasks(tasks);
    } catch (error) {
        console.error(error);
    }
}


document.getElementById('add-task-form').addEventListener('submit', async (e) => {
    e.preventDefault();
    const title = document.getElementById('task-title').value;

    try {
        const response = await fetch('http://localhost:8080/tasks', {
            method: 'POST',
            headers: { 
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify({ title })
        });

        if (!response.ok) throw new Error('Ошибка добавления задачи');

        document.getElementById('task-title').value = '';
        loadTasks();
    } catch (error) {
        console.error(error);
    }
});


function renderTasks(tasks) {
    const list = document.getElementById('tasks-list');
    list.innerHTML = tasks.map(task => `
        <li class="task-item">
            <span>${task.title}</span>
            <button onclick="deleteTask(${task.id})">Удалить</button>
        </li>
    `).join('');
}

window.deleteTask = async (id) => {
    try {
        const response = await fetch(`http://localhost:8080/tasks/${id}`, {
            method: 'DELETE',
            headers: { 'Authorization': token }
        });

        if (!response.ok) throw new Error('Ошибка удаления задачи');

        loadTasks();
    } catch (error) {
        console.error(error);
    }
};

function showTasksInterface() {
    document.getElementById('auth-form').style.display = 'none';
    document.getElementById('register-form').style.display = 'none';
    document.getElementById('tasks-interface').style.display = 'block';
}


if (token) {
    showTasksInterface();
    loadTasks();
}