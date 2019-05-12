// import React from "react"
// import {Query} from "react-apollo"
// import gql from "graphql-tag"

// const query = gql`
//     query Visitor {
//         User {
//             id, name
//         }
//     }
//     `

// export const Index: React.StatelessComponent = () => (
//     <div className="Index">
//         <h1>Diaries</h1>
//         <Query<ListDiaries> query={query}>
//             {result => {
//                 if (result.error) {
//                     return <p className="error">Error: {result.error.message}</p>                    
//                 }
//                 if (result.loading) {
//                     return <p className="loading">Loading</p>
//                 }
//                 const { data } = result;
//                 return <DiaryList diaries={data!.listDiaries} />;
//             }}
//         </Query>
//     </div>
// );