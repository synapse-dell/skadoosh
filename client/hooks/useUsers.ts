"use client"

import { useCallback, useEffect, useState } from "react"
import type { User } from "@/types";
import { usersApi } from "@/services/api";

export function useUsers() {
    const [users, setUsers] = useState<User[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    const reload = useCallback(async () => {
        try {
            setLoading(true);
            setError(null);
            const data = await usersApi.list();
            setUsers(data);
        } catch (error: any) {
            setError(
                error instanceof Error
                    ? error.message
                    : "Something went wrong"
            );
        } finally {
            setLoading(false);
        }
    }, []);

    useEffect(() => {
        void reload();
    }, [reload]);

    return {
        users,
        reload,
        loading,
        error
    }
}