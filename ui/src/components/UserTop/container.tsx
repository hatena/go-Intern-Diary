import React from "react"
import gql from "graphql-tag"

import {diaryListFragment} from "../DiaryList/diaryList"

import { UserTop } from "./userTop"

export const query = gql`
    query GetVisitor {
        visitor {
            ...DiaryListFragment
        }
    }
${diaryListFragment}
`

export const UserTopContainer: React.StatelessComponent = () => (
    <UserTop />
);