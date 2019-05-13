import React from "react";
import {Link} from "react-router-dom";
import gql from "graphql-tag";

interface DiaryListItemFragment {
    id: string;
    name: string;
}

interface DiaryListFragment {
    id: string;
    name: string
}

const diaryListItemFragment = gql`fragment DiaryListItemFragment on Diary {
    id
    name
}`;

interface DiaryListItemProps {
    diary: DiaryListItemFragment
}

export const DiaryListItem: React.StatelessComponent<DiaryListItemProps> = ({ diary }) => (
    <div className="DiaryListItem">
    <Link to={`/diaries/${diary.id}`}>{diary.name}</Link>
    <span> - </span>
  </div>
)

export const diaryListFragment = gql`fragment DiaryListFragment on Diary {
    id
    ...DiaryListItemFragment
}
${diaryListItemFragment}
`;

interface DiaryListProps {
    diaries: DiaryListFragment[]
}

export const DiaryList: React.StatelessComponent<DiaryListProps> = ({ diaries }) => (
    <ul className="DiaryList">
        {diaries.map(diary => (<li key={diary.id}><DiaryListItem diary={diary} /></li>))}
    </ul>
);
