import { TodoItem } from "../model/TodoItem";
import { ActionTypes, IInitStoreAction, IAddTodoAction, ICompleteTodoAction } from "./actionTypes";
import { fetch } from "cross-fetch";
import { Dispatch, Action } from "redux";
import { resolve } from "dns";

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

export const fetchTodos = () => {
    const url = 'http://localhost' + '/todo';
    return (dispatch: Dispatch<Action>) => {
        return fetch(url)
                .then(result => result.json())
                .then(mapTodoItems)
                .then(todoItems => {
                    return dispatch(initStoreAction(todoItems));
                });        
    }
}

function mapTodoItems(items: any[]): TodoItem[] {
    return items.map((item) => {
        return item as TodoItem;
    });
}