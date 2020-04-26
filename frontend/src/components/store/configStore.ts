import { createStore, applyMiddleware } from "redux";
import {composeWithDevTools} from "redux-devtools-extension/developmentOnly";
import thunk from "redux-thunk";
import { TodoItem } from "../model/TodoItem";
import { fetchTodosAsync, initStoreAction } from "../actions/actions";
import { rootReducer } from "../reducers/rootReducer";

export interface IState {
    todos: TodoItem[];
}

export const initStore = () => {
    return async (dispatch: any) => {
        const todos = await fetchTodosAsync<TodoItem[]>();
        return dispatch(initStoreAction(todos));
    };
};

export const configureStore = () => {
    if (process.env.NODE_ENV === "production") {
        return createStore(
            rootReducer,
            applyMiddleware(thunk),
        );
    } else {
        return createStore(
            rootReducer,
            composeWithDevTools(
                applyMiddleware(thunk),
            ),
        );
    }
};
