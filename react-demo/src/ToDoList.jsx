import React, { useState } from "react";

const ToDoList = () => {
    const [tasks, setTasks] = useState(["task1", "task2"]);
    const [newTask, setNewTask] = useState("");

    const handleInputChange = (e) => {
        setNewTask(e.target.value);
    };

    const addTask = () => {
        if (newTask.trim() !== "") {
            setTasks(t => [...t, newTask])
            setNewTask("")
        }
    };

    const deleteTask = (index) => {
        const updateTasks = tasks.filter((_, i) => i !== index)
        setTasks(updateTasks)
    };

    const moveTaskUp = (index) => {
        if (index > 0) {
            const updateTasks = [...tasks];
            [updateTasks[index], updateTasks[index - 1]] = [updateTasks[index - 1], updateTasks[index]]
            setTasks(updateTasks)
        }
    };
    const moveTaskDown = (index) => {
        if (index < tasks.length - 1) {
            const updateTasks = [...tasks];
            [updateTasks[index], updateTasks[index + 1]] = [updateTasks[index + 1], updateTasks[index]]
            setTasks(updateTasks)
        }
    };

    return (
        <div className="to-do-list">
            <h1>List of ToDo</h1>
            <input
                type="text"
                placeholder="Enter a task..."
                value={newTask}
                onChange={handleInputChange}
            />
            <button className="add-button" onClick={addTask}>Add</button>

            <ol>
                {tasks.map((task, index) =>
                    <li key={index} >
                        <span className="text"> {task}</span>
                        <button className="delete-button" onClick={() => deleteTask(index)}>Delete</button>
                        <button className="move-button" onClick={() => moveTaskUp(index)}>Move Up</button>
                        <button className="move-button" onClick={() => moveTaskDown(index)}>Move Down</button>
                    </li>

                )}
            </ol>
        </div>


    );
};

export default ToDoList;