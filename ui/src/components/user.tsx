import React from "react";
import {Link} from "react-router-dom";
import gql from "graphql-tag";

export interface UserType {
    id: string;
    name: string;
}

interface UserProps {
    user: UserType
}

export const User: React.StatelessComponent<UserProps> = ({ user }) => (
    <div className="user">
        id: {user.id} name: {user.name}
    </div>
);