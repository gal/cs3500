import { doRequest, NoData } from ".";
import { useAuth } from "../store/auth";
import { Project } from "./types";

export const GetProjectByID = async (projID: number): Promise<Project> => {
  const accessToken = useAuth((state) => state.accessToken);
  const [resp, error] = await doRequest({
    method: "GET",
    url: `/api/projects/${projID}`,
    headers: {
      "Content-Type": "application/json",
      "Authorization": `Bearer ${accessToken}`,
    },
  });

  if (error) {
    throw error;
  }

  return (resp?.data as Project) ?? NoData;
};

export const GetProjectsByOwnerID = async (
  userID: number
): Promise<Project[]> => {
  const [resp, error] = await doRequest({
    method: "GET",
    url: `/api/user/${userID}/projects`,
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (error) {
    throw error;
  }

  return (resp?.data as Project[]) ?? NoData;
};

export const GetProjectsByRequiredSkill = async (
  tagID: number
): Promise<Project[]> => {
  const [resp, error] = await doRequest({
    method: "GET",
    url: `/api/tag/${tagID}/projects`,
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (error) {
    throw error;
  }

  return (resp?.data as Project[]) ?? NoData;
};

export const GetProjectsByPreferredSkill = async (
  tagID: number
): Promise<Project[]> => {
  const [resp, error] = await doRequest({
    method: "GET",
    url: `/api/tag/${tagID}/projects`,
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (error) {
    throw error;
  }

  return (resp?.data as Project[]) ?? NoData;
};