import React from "react"
import {Mutation} from "react-apollo"
import gql from "graphql-tag"
import { MutationUpdaterFn } from "apollo-client";

import { query } from "../UserTop/container"
import{ GetVisitor } from "../UserTop/__generated__/GetVisitor"
import { DeleteDiary } from "./__generated__/DeleteDiary"
import { DeleteDiaryForm } from "./deleteDiary";

export const deleteDiary = gql`
    mutation DeleteDiary($diaryId: ID!) {
        deleteDiary(diaryId: $diaryId)
    }
`;

export const updateDiary: (diaryId: string) => MutationUpdaterFn<DeleteDiary> = (diaryId) => (cache, result) => {
    const visitor  = cache.readQuery<GetVisitor>({ query });
    const { data } = result
    if (visitor && data) {
        const diaries = [...visitor.visitor.diaries].filter(diary => diary.id !== diaryId);
        const newVisitor = {
            visitor: {
                ...visitor.visitor,
                diaries,
            }
        };
        cache.writeQuery({query, data: newVisitor})
    }
}

interface DeleteDiaryContainerProps {
    diaryId: string;
}

export const DeleteDiaryFormContainer: React.StatelessComponent<DeleteDiaryContainerProps> = ({ diaryId }) => (
    <DeleteDiaryForm diaryId={diaryId}/>
)