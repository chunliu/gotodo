import { combineReducers } from "redux";
import { todos } from "./todosReducer";

export const rootReducer = combineReducers({
    todos,
});
