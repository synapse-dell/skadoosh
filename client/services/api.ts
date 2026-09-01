import type { User } from "@/types";

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

async function request<T>(
    endpoint: string,
    options?: RequestInit
): Promise<T> {
    const response = await fetch(`${API_BASE_URL}${endpoint}`, {
        ...options,
        headers: {
            "Content-Type": "application/json",
            ...options?.headers,
        },
    })
    if (!response.ok) {
        throw new Error("Something Went wrong");
    }
    return response.json();
}

export const usersApi = {
    async list(): Promise<User[]> {
        const response = await request<{
            message: string,
            Users: User[]
        }>("/");

        return response.Users;
    },
    async create(input: Omit<User, "id">): Promise<User> {
        const response = await request<{ user: User }>(
            "/create",
            {
                method: "POST",
                body: JSON.stringify(input),
            }
        );

        return response.user;
    },
    async update(
        id: number,
        input: Omit<User, "id">
    ): Promise<User> {
        const response = await request<{ user: User }>(
            `/update/${id}`,
            {
                method: "PUT",
                body: JSON.stringify(input)
            }
        );
        return response.user;
    },
    async delete(
        id: number,
    ): Promise<User> {
        const response = await request<{ user: User }>(
            `/delete/${id}`,
            {
                method: "DELETE",
            }
        )
        return response.user;
    }
}

