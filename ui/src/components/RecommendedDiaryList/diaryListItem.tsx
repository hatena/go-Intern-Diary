import React from "react"
import {Link} from "react-router-dom";
import {RecommendedDiary} from "./container"

interface DiaryListItemProps {
    diary: RecommendedDiary;
}

export const DiaryListItem: React.StatelessComponent<DiaryListItemProps> = ({diary}) => (
    <div key={diary.id}>
        <Link to={`/diaries/${diary.id}`}>
            <h4><span>{diary.userName}</span>さんの{diary.diaryName}</h4>
        </Link>
    </div>
)