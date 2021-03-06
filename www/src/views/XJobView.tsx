import { useQuery } from "@apollo/react-hooks";
import { ApolloError } from "apollo-client/errors/ApolloError";
import gql from "graphql-tag";
import * as React from "react";
import { useState } from "react";
import { useParams } from "react-router-dom";
import { Header, Icon } from "semantic-ui-react";
import { XJobHeader } from "../components/job";
import { XCardGroup } from "../components/layout";
import { XErrorMessage, XLoadingMessage } from "../components/messages";
import { XTaskCard, XTaskContent } from "../components/task";
import { XTaskCardDisplayType } from "../components/task/XTaskCard";
import { Job, Tag, Task } from "../graphql/models";

export const JOB_QUERY = gql`
  query Job($id: ID!) {
    job(id: $id) {
      id
      name
      content
      tags {
        id
        name
      }
      tasks {
        id
        queueTime
        claimTime
        execStartTime
        execStopTime
        content
        output
        error
        sessionID

        target {
          id
          name
        }
        job {
          id
          name
          tags {
            id
            name
          }
        }
      }
    }
  }
`;

type JobQuery = {
  job: Job;
};

const XJobView = () => {
  let { id } = useParams();
  const [error, setError] = useState<ApolloError>(null);

  const [name, setName] = useState<string>("");
  const [content, setContent] = useState<string>("");
  const [tags, setTags] = useState<Tag[]>([]);
  const [tasks, setTasks] = useState<Task[]>([]);

  const { called, loading } = useQuery<JobQuery>(JOB_QUERY, {
    variables: { id },
    pollInterval: 5000,
    onCompleted: data => {
      setError(null);
      if (!data || !data.job) {
        data = { job: { id: "", name: "", content: "", tags: [], tasks: [] } };
      }

      setName(data.job.name || "");
      setContent(data.job.content || "");
      setTags(data.job.tags || []);
      setTasks(data.job.tasks || []);
    },
    onError: err => setError(err)
  });

  const showCards = () => {
    if (!tasks || tasks.length < 1) {
      return (
        // TODO: Better styling
        <h1>No tasks found!</h1>
      );
    }
    return (
      <XCardGroup>
        {tasks.map(task => (
          <XTaskCard
            key={task.id}
            display={XTaskCardDisplayType.TARGET}
            task={task}
          />
        ))}
      </XCardGroup>
    );
  };

  return (
    <React.Fragment>
      <XJobHeader name={name} tags={tags} />

      <XErrorMessage title="Error Loading Job" err={error} />
      <XLoadingMessage
        title="Loading Job"
        msg="Fetching job information..."
        hidden={called && !loading}
      />

      <XTaskContent content={content} />

      <Header size="large" block inverted>
        <Icon name="tasks" />
        <Header.Content>Tasks</Header.Content>
      </Header>

      {showCards()}
    </React.Fragment>
  );
};

export default XJobView;
