import React from "react";
import {Link} from "react-router-dom";
import gql from "graphql-tag";
import { DiaryListFragment } from "./__generated__/DiaryListFragment";
import { DiaryListItemFragment } from "./__generated__/DiaryListItemFragment"


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

interface DiaryListItemProps {
    diary: DiaryListItemFragment
    deleteDiary?: () => void
}

export const DiaryListItem: React.StatelessComponent<DiaryListItemProps> = ({ diary, deleteDiary }) => (
    <div className="DiaryListItem">
        <div>
            <Link to={`/diaries/${diary.id}`}>{diary.name}</Link>
        </div>
        <div>
            {deleteDiary && <button onClick={deleteDiary}>Delete</button>}
        </div>
    </div>
)

interface DiaryListProps {
    user: DiaryListFragment
    deleteDiary?: (diaryId: string) => void
}

export const DiaryList: React.StatelessComponent<DiaryListProps> = ({ user, deleteDiary }) => (
    <div className="DiaryList">
        <h1>{user.name}'s Diaries</h1>
        <ul>
            {user.diaries.map(diary => (
                <li key={diary.id}>
                    <DiaryListItem 
                    diary={diary} 
                    deleteDiary={deleteDiary ? () => {deleteDiary(diary.id) }: undefined} />
                </li>
            ))}
        </ul>
    </div>
);
