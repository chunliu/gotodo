import { createStore, applyMiddleware, Dispatch, Action } from "redux";
import {composeWithDevTools} from "redux-devtools-extension/developmentOnly";
import thunk from "redux-thunk";
import { TodoItem } from "../model/TodoItem";
import { fetchTodos } from "../actions/actions";
import { rootReducer } from "../reducers/rootReducer";

export interface IState {
    todos: TodoItem[];
}

export const initStore = () => {
    return (dispatch: any) => {
        return dispatch(fetchTodos());
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
