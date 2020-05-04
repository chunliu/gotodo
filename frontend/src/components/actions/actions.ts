import { TodoItem } from "../model/TodoItem";
import { ActionTypes, IInitStoreAction, IAddTodoAction, ICompleteTodoAction } from "./actionTypes";
import { fetch } from "cross-fetch";

const baseUrl = (process.env.NODE_ENV === "development") ? "http://localhost" : "";

export const initStoreAction = (todos: TodoItem[]): IInitStoreAction => {
    return {type: ActionTypes.INIT_STORE, todos};
};

export const addTodoAction = (todo: TodoItem): IAddTodoAction => {
    return {type: ActionTypes.ADD_TODO_ITEM, todo};
};

export function addTodo(todo: TodoItem) {
    return async (dispatch: any) => {
        const resp = await addTodoAsync(todo);
        if(resp.ok) {
            dispatch(addTodoAction(todo));
        }
    }
}

export const completeTodoAction = (todo: TodoItem): ICompleteTodoAction => {
    return {type: ActionTypes.COMPLETE_TODO_ITEM, todo};
};

export function completeTodo(todo: TodoItem) {
    return async (dispatch: any) => {
        const resp = await updateTodoAsync(todo);
        if(resp.ok) {
            dispatch(completeTodoAction(todo));
        }
    }
}

export const actionCreators = {
    addTodo,
    completeTodo,
};

export async function fetchTodosAsync<T>(): Promise<T> {
    const url = baseUrl + "/todo";
    const resp = await fetch(url);
    const data = await resp.json();
    
    return data;
}

export async function addTodoAsync(todo: TodoItem): Promise<Response> {
    const url = baseUrl + "/todo";
    const resp = await fetch(url, {
            method: "POST",
            body: JSON.stringify(todo),
        });
    return resp;
}

export async function updateTodoAsync(todo: TodoItem): Promise<Response> {
    const url = baseUrl + "/todo/" + todo.id;
    const resp = await fetch(url, {
            method: "PUT",
            body: JSON.stringify(todo),
        });
    return resp;
}