<!DOCTYPE html>
<html>
<body>
    <h1>Мои задачи</h1>
    <div id="tasks"></div>
    
    <script>
        fetch('http://localhost:8080/tasks')
            .then(res => res.json())
            .then(tasks => {
                document.getElementById('tasks').innerHTML = 
                    tasks.map(t => `<div>${t.title}</div>`).join('')
            });
    </script>
</body>
</html>