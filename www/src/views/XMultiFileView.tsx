import { useQuery } from "@apollo/react-hooks";
import gql from "graphql-tag";
import * as React from "react";
import { Loader } from "semantic-ui-react";
import { XFileCard, XNoFilesFound } from "../components/file";
import { XCardGroup } from "../components/layout";
import { XErrorMessage } from "../components/messages";
import { File } from "../graphql/models";

export const MULTI_FILE_QUERY = gql`
  {
    files {
      id
      name
      contentType
      size
      creationTime
      lastModifiedTime

      links {
        id
        alias
        clicks
        expirationTime
      }
    }
  }
`;

type MultiFile = {
  files: File[];
};

const XMultiFileView = () => {
  const { called, loading, error, data } = useQuery<MultiFile>(
    MULTI_FILE_QUERY,
    {
      pollInterval: 5000
    }
  );

  const showCards = () => {
    if (!data || !data.files || data.files.length < 1) {
      return <XNoFilesFound />;
    }
    return (
      <XCardGroup>
        {data.files.map(file => (
          <XFileCard key={file.id} {...file} />
        ))}
      </XCardGroup>
    );
  };

  return (
    <React.Fragment>
      <Loader disabled={!called || !loading} />
      <XErrorMessage title="Error Loading Files" err={error} />
      {showCards()}
    </React.Fragment>
  );
};

// XTargetCardGroup.propTypes = {
//     targets: PropTypes.arrayOf(PropTypes.shape({
//         id: PropTypes.number.isRequired,
//         name: PropTypes.string.isRequired,
//         tags: PropTypes.arrayOf(PropTypes.string),
//     })).isRequired,
// };

export default XMultiFileView;
