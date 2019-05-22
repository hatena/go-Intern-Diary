import React from "react";
import gql from "graphql-tag";
import { DiaryListFragment } from "./__generated__/DiaryListFragment";
import { DiaryListItem } from "../DiaryListItem/diaryListItem"



export const diaryListItemFragment = gql`fragment DiaryListItemFragment on Diary {
    id
    name
}`;

export const diaryListFragment = gql`fragment DiaryListFragment on User {
    name
    diaries {
        ...DiaryListItemFragment
    }
}
${diaryListItemFragment}
`;


interface DiaryListProps {
    user: DiaryListFragment
}

export const DiaryList: React.StatelessComponent<DiaryListProps> = ({ user }) => (
    <div className="DiaryList">
        <h1>{user.name}'s Diaries</h1>
        <ul>
            {user.diaries.map(diary => (
                <li key={diary.id}>
                    <DiaryListItem diary={diary} />
                </li>
            ))}
        </ul>
    </div>
);
