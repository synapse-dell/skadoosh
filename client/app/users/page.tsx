"use client";

import { useState } from "react";
import { usersApi } from "@/services/api";
import { useUsers } from "@/hooks/useUsers";

export default function UsersPage() {
  const { users, loading, error, reload } = useUsers();

  const [name, setName] = useState("");
  const [age, setAge] = useState("");

  async function handleCreate(e: React.FormEvent) {
    e.preventDefault();

    await usersApi.create({
      name,
      age: Number(age),
    });

    setName("");
    setAge("");

    await reload();
  }

  async function handleDelete(id: number) {
    await usersApi.delete(id);
    await reload();
  }

  async function handleUpdate(user: {
    id: number;
    name: string;
    age: number;
  }) {
    await usersApi.update(user.id, {
      name: user.name,
      age: user.age + 1,
    });

    await reload();
  }

  if (loading) {
    return <p>Loading...</p>;
  }

  if (error) {
    return (
      <div>
        <p>{error}</p>
        <button onClick={reload}>Retry</button>
      </div>
    );
  }

  return (
    <main>
      <h1>Users</h1>

      <form onSubmit={handleCreate}>
        <input
          value={name}
          onChange={(e) => setName(e.target.value)}
          placeholder="Name"
        />

        <input
          type="number"
          value={age}
          onChange={(e) => setAge(e.target.value)}
          placeholder="Age"
        />

        <button type="submit">
          Add User
        </button>
      </form>

      <hr />

      {users.map((user) => (
        <div key={user.id}>
          <span>
            {user.name} - {user.age}
          </span>

          <button onClick={() => handleUpdate(user)}>
            +1 Age
          </button>

          <button onClick={() => handleDelete(user.id)}>
            Delete
          </button>
        </div>
      ))}
    </main>
  );
}