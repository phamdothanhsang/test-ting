export interface Kols {
  KolID: number;
  UserProfileID: number;
  Language: string;
  Education: string;
  ExpectedSalary: number;
  ExpectedSalaryEnable: boolean;
  ChannelSettingTypeID: number;
  IDFrontURL: string;
  IDBackURL: string;
  PortraitURL: string;
  RewardID: number;
  PaymentMethodID: number;
  TestimonialsID: number;
  VerificationStatus: string;
  Enabled: boolean;
  ActiveDate: Date;
  Active: boolean;
  CreatedBy: string;
  CreatedDate: Date;
  ModifiedBy: string;
  ModifiedDate: Date;
  IsRemove: boolean;
  IsOnBoarding: boolean;
  Code: string;
  PortraitRightURL: string;
  PortraitLeftURL: string;
  LivenessStatus: string;
}

export const transformKol = (item: any) => {
  const transformedData: Kols = {
    KolID: item.kolID,
    UserProfileID: item.userProfileID,
    Language: item.language,
    Education: item.education,
    ExpectedSalary: item.expectedSalary,
    ExpectedSalaryEnable: item.expectedSalaryEnable,
    ChannelSettingTypeID: item.channelSettingTypeID,
    IDFrontURL: item.iDFrontURL,
    IDBackURL: item.iDBackURL,
    PortraitURL: item.portraitURL,
    RewardID: item.rewardID,
    PaymentMethodID: item.paymentMethodID,
    TestimonialsID: item.testimonialsID,
    VerificationStatus: item.verificationStatus,
    Enabled: item.enabled,
    ActiveDate: new Date(item.activeDate),
    Active: item.active,
    CreatedBy: item.createdBy,
    CreatedDate: new Date(item.createdDate),
    ModifiedBy: item.modifiedBy,
    ModifiedDate: new Date(item.modifiedDate),
    IsRemove: item.isRemove,
    IsOnBoarding: item.isOnBoarding,
    Code: item.code,
    PortraitRightURL: item.portraitRightURL,
    PortraitLeftURL: item.portraitLeftURL,
    LivenessStatus: item.livenessStatus,
  };

  return transformedData;
};
