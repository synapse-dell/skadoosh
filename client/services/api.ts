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
    async List(): Promise<User[]> {
        const response = await request<{
            message: string,
            Users: User[]
        }>("/users");

        return response.Users;
    }
}

