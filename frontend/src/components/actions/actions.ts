import { TodoItem } from "../model/TodoItem";
import { ActionTypes, IInitStoreAction, IAddTodoAction, ICompleteTodoAction } from "./actionTypes";
import { fetch } from "cross-fetch";

export const initStoreAction = (todos: TodoItem[]): IInitStoreAction => {
    return {type: ActionTypes.INIT_STORE, todos};
};

export const addTodoAction = (todo: TodoItem): IAddTodoAction => {
    return {type: ActionTypes.ADD_TODO_ITEM, todo};
};

export const completeTodoAction = (todo: TodoItem): ICompleteTodoAction => {
    return {type: ActionTypes.COMPLETE_TODO_ITEM, todo};
};

export const actionCreators = {
    addTodoAction,
    completeTodoAction,
};

export async function fetchTodosAsync<T>(): Promise<T> {
    const url = 'http://localhost' + '/todo';
    const resp = await fetch(url);
    const data = await resp.json();
    
    return data;
}

export async function addTodoAsync(todo: TodoItem): Promise<Response> {
    const url = 'http://localhost' + '/todo';
    const resp = await fetch(url, {
            method: "POST",
            body: JSON.stringify(todo),
        });
    return resp;
}