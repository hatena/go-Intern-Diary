import React from "react"
import {Mutation} from "react-apollo"
import {deleteDiary, updateDiary} from "./container"
import { DeleteDiary, DeleteDiaryVariables } from "./__generated__/DeleteDiary"


interface DeleteDiaryProps {
    diaryId: string;
}

export const DeleteDiaryForm: React.StatelessComponent<DeleteDiaryProps> = ({ diaryId }) => (
    <Mutation<DeleteDiary, DeleteDiaryVariables> mutation={deleteDiary}>
        {(deleteDiary) => {
            return (
                <div>
                    {deleteDiary && <button 
                    onClick={ () => deleteDiary({ variables: {diaryId: diaryId}, update: updateDiary(diaryId)})}>Delete</button>}
                </div>
            )
    }}
    </Mutation>
)