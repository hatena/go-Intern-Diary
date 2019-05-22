import React from "react"
import {Link} from "react-router-dom";
import { DeleteDiaryForm} from "./deleteDiary"
import { DiaryListItemFragment } from "../DiaryList/__generated__/DiaryListItemFragment"

interface DiaryListItemProps {
    diary: DiaryListItemFragment
}

export const DiaryListItem: React.StatelessComponent<DiaryListItemProps> = ({ diary }) => (
    <div className="DiaryListItem">
        <div>
            <Link to={`/diaries/${diary.id}`}>{diary.name}</Link>
        </div>
        <DeleteDiaryForm diaryId={diary.id} />
    </div>       
)